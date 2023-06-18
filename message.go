package youdu

const (
	apiMsgSend = "/cgi/msg/send"
)

const (
	MsgTypeText  = "text"
	MsgTypeImage = "image"
	MsgTypeFile  = "file"
)

type (
	MessageRequest interface{}

	TextMessageRequest struct {
		ToUser  string `json:"toUser,omitempty"`
		ToDept  string `json:"toDept,omitempty"`
		MsgType string `json:"msgType"`
		Text    struct {
			Content string `json:"content"`
		} `json:"text"`
	}

	MessageResponse struct {
		ErrCode int    `json:"errcode"`
		ErrMsg  string `json:"errmsg"`
	}
)

var _ MessageRequest = (*TextMessageRequest)(nil)

func (c *Client) SendMessage(request MessageRequest) (response MessageResponse, err error) {
	req, err := c.newRequestWithToken("POST", apiMsgSend, request)
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	return
}

func (c *Client) SendTextMessage(request TextMessageRequest) (MessageResponse, error) {
	return c.SendMessage(request)
}
