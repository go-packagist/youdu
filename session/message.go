package session

type Message interface {
}

var _ Message = (*TextMessage)(nil)
var _ Message = (*ImageMessage)(nil)
var _ Message = (*VoiceMessage)(nil)
var _ Message = (*VideoMessage)(nil)

const (
	MsgTypeText  = "text"
	MsgTypeImage = "image"
	MsgTypeVoice = "voice"
	MsgTypeVideo = "video"
)

type TextItem struct {
	Content string `json:"content"`
}

type TextMessage struct {
	SessionId string    `json:"sessionId,omitempty"`
	Receiver  string    `json:"receiver,omitempty"`
	Sender    string    `json:"sender"`
	MsgType   string    `json:"msgType"`
	Text      *TextItem `json:"text"`
}

type MediaItem struct {
	MediaId string `json:"mediaId"`
}

type ImageMessage struct {
	SessionId string     `json:"sessionId,omitempty"`
	Receiver  string     `json:"receiver,omitempty"`
	Sender    string     `json:"sender"`
	MsgType   string     `json:"msgType"`
	Image     *MediaItem `json:"image"`
}

type FileMessage struct {
	SessionId string     `json:"sessionId,omitempty"`
	Receiver  string     `json:"receiver,omitempty"`
	Sender    string     `json:"sender"`
	MsgType   string     `json:"msgType"`
	File      *MediaItem `json:"file"`
}

type VoiceMessage struct {
	SessionId string     `json:"sessionId,omitempty"`
	Receiver  string     `json:"receiver,omitempty"`
	Sender    string     `json:"sender"`
	MsgType   string     `json:"msgType"`
	Voice     *MediaItem `json:"voice"`
}

type VideoMessage struct {
	SessionId string     `json:"sessionId,omitempty"`
	Receiver  string     `json:"receiver,omitempty"`
	Sender    string     `json:"sender"`
	MsgType   string     `json:"msgType"`
	Video     *MediaItem `json:"video"`
}
