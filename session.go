package youdu

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

const (
	sessionCreateUrl = "/cgi/session/create"
	sessionGetUrl    = "/cgi/session/get"
	sessionUpdateUrl = "/cgi/session/update"
)

type Session struct {
	config *Config
}

func NewSession(config *Config) *Session {
	return &Session{
		config: config,
	}
}

type SessionInfo struct {
	SessionId  string   `json:"sessionId"`
	Type       string   `json:"type"`
	Owner      string   `json:"owner"`
	Title      string   `json:"title"`
	Version    int      `json:"version"`
	Member     []string `json:"member"`
	LastMsgId  int      `json:"lastMsgId"`
	ActiveTime int      `json:"activeTime"`
}

// CreateSession 创建一个会话
// members 第一个默认为创建者
func (s *Session) CreateSession(title string, members []string) (*SessionInfo, error) {
	if len(members) < 3 {
		return nil, errors.New("members must be at least 3")
	}

	accessToken, err := s.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"title":   title,
		"creator": members[0],
		"member":  members,
		"type":    "multi",
	})
	if err != nil {
		return nil, err
	}

	encrypt, err := s.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return nil, err
	}

	resp, err := s.config.GetHttp().Post(sessionCreateUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   s.config.AppId,
		"buin":    s.config.Buin,
		"encrypt": encrypt,
	})

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return nil, err
	}

	decrypt, err := s.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	var v *SessionInfo
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v, nil
}

// GetSession 获取会话信息
func (s *Session) GetSession(sessionId string) (*SessionInfo, error) {
	accessToken, err := s.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	resp, err := s.config.GetHttp().Get(sessionGetUrl+"?accessToken="+accessToken, map[string]string{
		"sessionId": sessionId,
	})

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return nil, err
	}

	decrypt, err := s.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	var v *SessionInfo
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v, nil
}

// UpdateSession 更新会话信息
func (s *Session) UpdateSession(sessionId, opUser, title string, addMembers, delMembers []string) (*SessionInfo, error) {
	accessToken, err := s.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"sessionId": sessionId,
		"opUser":    opUser,
		"title":     title,
		"addMember": addMembers,
		"delMember": delMembers,
	})
	if err != nil {
		return nil, err
	}

	encrypt, err := s.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return nil, err
	}

	resp, err := s.config.GetHttp().Post(sessionUpdateUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   s.config.AppId,
		"buin":    s.config.Buin,
		"encrypt": encrypt,
	})

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return nil, err
	}

	decrypt, err := s.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	fmt.Println(decrypt)

	var v *SessionInfo
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v, nil
}
