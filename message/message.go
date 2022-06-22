package message

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
var _ Message = (*ImageMessage)(nil)
var _ Message = (*FileMessage)(nil)
var _ Message = (*MpNewsMessage)(nil)
var _ Message = (*LinkMessage)(nil)
var _ Message = (*ExLinkMessage)(nil)
var _ Message = (*SmsMessage)(nil)
var _ Message = (*MailMessage)(nil)

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

type SysMessage struct {
	ToUser string `json:"toUser"`
	ToDept string `json:"toDept"`
	ToAll  struct {
		OnliyOnline bool `json:"onliyOnline"`
	} `json:"toAll"`
	MsgType string `json:"msgType"`
	SysMsg  struct {
		Title       string `json:"title"`
		PopDuration int    `json:"popDuration"`
		Msg         []struct {
			Text struct {
				Content string `json:"content"`
			} `json:"text,omitempty"`
			Link struct {
				Title  string `json:"title"`
				Url    string `json:"url"`
				Action int    `json:"action"`
			} `json:"link,omitempty"`
		} `json:"msg"`
	} `json:"sysMsg"`
}

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

type PopWindowItem struct {
	Url      string `json:"url"`
	Tip      string `json:"tip"`
	Title    string `json:"title"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	Duration int    `json:"duration"`
	Position int    `json:"position"`
	NoticeId string `json:"notice_id"`
	PopMode  int    `json:"pop_mode"`
}

type PopWindowMessage struct {
	ToUser    string         `json:"toUser"`
	ToDept    string         `json:"toDept"`
	PopWindow *PopWindowItem `json:"popWindow"`
}
