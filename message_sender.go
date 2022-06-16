package youdu

import (
	"encoding/json"
	"errors"
	"strconv"
)

const msgSendUrl = "/cgi/msg/send"

// see:https://youdu.im/doc/api/c01_00003.html#_7
const (
	MsgTypeText   = "text"
	MsgTypeImage  = "image"
	MsgTypeFile   = "file"
	MsgTypeMpNews = "mpnews"
	MsgTypeAudio  = "audio"
	MsgTypeVideo  = "video"
	MsgTypeLink   = "link"
)

type Message interface {
}

var _ Message = (*TextMessage)(nil)

type TextItem struct {
	Content string `json:"content"`
}

type TextMessage struct {
	ToUser  string    `json:"toUser"`
	ToDept  string    `json:"toDept"`
	MsgType string    `json:"msgType"`
	Text    *TextItem `json:"text"`
}

type MediaItem struct {
	MediaId string `json:"media_id"`
}

type ImageMessage struct {
	ToUser  string     `json:"toUser"`
	ToDept  string     `json:"toDept"`
	MsgType string     `json:"msgType"`
	Image   *MediaItem `json:"image"`
}

type FileMessage struct {
	ToUser  string     `json:"toUser"`
	ToDept  string     `json:"toDept"`
	MsgType string     `json:"msgType"`
	File    *MediaItem `json:"file"`
}

type MpNewsItem struct {
	Title     string `json:"title"`
	MediaId   string `json:"media_id"`
	Content   string `json:"content"`
	Digest    string `json:"digest"`
	ShowFront int    `json:"show_front"`
}

type MpNewsMessage struct {
	ToUser  string        `json:"toUser"`
	ToDept  string        `json:"toDept"`
	MsgType string        `json:"msgType"`
	MpNews  []*MpNewsItem `json:"mpNews"`
}

type LinkItem struct {
	Title  string `json:"title"`
	Url    string `json:"url"`
	Action int    `json:"action"`
}

type LinkMessage struct {
	ToUser  string    `json:"toUser"`
	ToDept  string    `json:"toDept"`
	MsgType string    `json:"msgType"`
	Link    *LinkItem `json:"link"`
}

type ExLinkItem struct {
	Title   string `json:"title"`
	Url     string `json:"url"`
	Digest  string `json:"digest"`
	MediaId string `json:"media_id"`
}

type ExLinkMessage struct {
	ToUser  string        `json:"toUser"`
	ToDept  string        `json:"toDept"`
	MsgType string        `json:"msgType"`
	ExLink  *[]ExLinkItem `json:"exlink"`
}

// type SysMessage struct {
// 	ToUser string `json:"toUser"`
// 	ToDept string `json:"toDept"`
// 	ToAll  struct {
// 		OnlyOnline bool `json:"only_online"`
// 	} `json:"toAll"`
// 	MsgType string `json:"msgType"`
// 	SysMsg
// }

type SmsMessage struct {
	ToUser  string `json:"toUser"`
	ToDept  string `json:"toDept"`
	MsgType string `json:"msgType"`
	Sms     *struct {
		From    string `json:"from"`
		Content string `json:"content"`
	} `json:"sms"`
}

type MailMessage struct {
	ToUser  string `json:"toUser"`
	ToEmail string `json:"toEmail"`
	MsgType string `json:"msgType"`
	Mail    *struct {
		Action      string `json:"action"`
		Subject     string `json:"subject"`
		FromUser    string `json:"fromUser"`
		FromEmail   string `json:"fromEmail"`
		Time        int    `json:"time"`
		Link        string `json:"link"`
		UnreadCount int    `json:"unreadCount"`
	}
}

type MessageSender struct {
	config *Config
}

func NewMessageSender(config *Config) *MessageSender {
	return &MessageSender{
		config: config,
	}
}

func (m *MessageSender) Send(message Message) error {
	accessToken, err := m.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return err
	}

	messageJson, err := json.Marshal(message)
	if err != nil {
		return err
	}

	encrypt, err := m.config.GetEncryptor().Encrypt(string(messageJson))
	if err != nil {
		return err
	}

	resp, err := m.config.GetHttp().Post(msgSendUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   m.config.AppId,
		"buin":    m.config.Buin,
		"encrypt": encrypt,
	})
	if err != nil {
		return err
	}

	if !resp.IsSuccess() {
		return errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return err
	}

	if jsonRet["errcode"].(float64) != 0 {
		return errors.New(jsonRet["errmsg"].(string))
	}

	return nil
}

func (m *MessageSender) SendText(toUser, content string, toDept ...string) error {
	if len(toDept) == 0 {
		toDept = []string{""}
	}

	return m.Send(&TextMessage{
		ToUser:  toUser,
		ToDept:  toDept[0],
		MsgType: MsgTypeText,
		Text: &TextItem{
			Content: content,
		},
	})
}
