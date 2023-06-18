package youdu

import (
	"strconv"
	"time"
)

const (
	apiGetToken = "/cgi/gettoken"
)

type Token struct {
	AccessToken string
	Expire      time.Time
}

type GetTokenRequest struct {
	Buin    int    `json:"buin"`
	AppId   string `json:"appId"`
	Encrypt string `json:"encrypt"`
}

type GetTokenResponse struct {
	AccessToken string `json:"AccessToken"`
	ExpireIn    int    `json:"expireIn"`
}

func (c *Client) GetToken() (t Token, err error) {
	if c.token != nil && c.token.AccessToken != "" && c.token.Expire.After(time.Now()) {
		return *c.token, nil
	}

	resp, err := c.GetAccessToken()
	if err != nil {
		return
	}

	c.token = &Token{
		AccessToken: resp.AccessToken,
		Expire:      time.Now().Add(time.Duration(resp.ExpireIn) * time.Second),
	}

	return *c.token, nil
}

func (c *Client) GetAccessToken() (response GetTokenResponse, err error) {
	encrypt, err := c.encryptor.Encrypt(strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		return
	}

	req, err := c.newRequest("POST", apiGetToken, GetTokenRequest{
		Buin:    c.config.Buin,
		AppId:   c.config.AppId,
		Encrypt: encrypt,
	})
	if err != nil {
		return
	}

	err = c.sendRequest(req, &response)

	return
}
