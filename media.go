package youdu

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"os"
	"strconv"
	"time"
)

const (
	mediaUploadUrl = "/cgi/media/upload"
	mediaSearchUrl = "/cgi/media/search"
)

const (
	MediaTypeImage = "image"
	MediaTypeFile  = "file"
	MediaTypeVoice = "voice"
	MediaTypeVideo = "video"
)

type Media struct {
	config *Config
}

func NewMedia(config *Config) *Media {
	return &Media{
		config: config,
	}
}

func (m *Media) Upload(fileType string, filePath string) (string, error) {
	accessToken, err := m.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return "", err
	}

	// encrypt
	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return "", err
	}

	fileByte, err := json.Marshal(map[string]interface{}{
		"type": fileType,
		"name": fileInfo.Name(),
	})
	if err != nil {
		return "", err
	}
	encrypt, err := m.config.GetEncryptor().Encrypt(string(fileByte))
	if err != nil {
		return "", err
	}

	// 加密文件
	contentByte, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	fileEncrypt, err := m.config.GetEncryptor().Encrypt(string(contentByte))
	if err != nil {
		return "", err
	}
	tmpFile := "/tmp/youdu-" + fileInfo.Name() + time.Now().Format("20060102150405") + ".tmp"
	defer os.Remove(tmpFile)
	if err := os.WriteFile(tmpFile, []byte(fileEncrypt), 0644); err != nil {
		return "", err
	}

	resp, err := m.config.GetHttp().Post(
		mediaUploadUrl+"?accessToken="+accessToken,
		map[string]interface{}{},
		func(request *resty.Request) {
			request.SetHeader("Content-Type", "multipart/form-data")
			request.SetFormData(map[string]string{
				"appId":   m.config.AppId,
				"buin":    strconv.Itoa(m.config.Buin),
				"encrypt": encrypt,
			})
			request.SetFile("file", tmpFile)
		})

	if err != nil {
		return "", err
	}

	if !resp.IsSuccess() {
		return "", errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return "", err
	}

	decrypt, err := m.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return "", err
	}

	var v map[string]string
	if err := decrypt.Unmarshal(&v); err != nil {
		return "", err
	}

	return v["mediaId"], err
}

type MediaInfo struct {
	Name string `json:"name"`
	Size int    `json:"size"`
}

func (m *Media) Search(mediaId string) (i MediaInfo, err error) {
	accessToken, err := m.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return
	}

	bodyJson, err := json.Marshal(map[string]string{
		"mediaId": mediaId,
	})
	if err != nil {
		return
	}

	encrypt, err := m.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return
	}

	resp, err := m.config.GetHttp().Post(mediaSearchUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   m.config.AppId,
		"buin":    m.config.Buin,
		"encrypt": encrypt,
	})

	if err != nil {
		return
	}

	if !resp.IsSuccess() {
		err = errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
		return
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return
	}

	if jsonRet["errcode"].(float64) != 0 {
		err = errors.New(jsonRet["errmsg"].(string))
		return
	}

	decrypt, err := m.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return
	}

	var v MediaInfo
	if err = decrypt.Unmarshal(&v); err != nil {
		return
	}

	return v, nil
}
