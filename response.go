package youdu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Response struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
	Encrypt string `json:"encrypt"`
}

func (c *Client) decodeResponse(body io.Reader, v interface{}) error {
	if v == nil {
		return nil
	}

	if result, ok := v.(*string); ok {
		return c.decodeString(body, result)
	}

	var (
		resp    Response
		decoder = json.NewDecoder(body)
	)
	if err := decoder.Decode(&resp); err != nil {
		return decoder.Decode(v)
	}

	return c.decodeEncryptResponse(resp, v)
}

func (c *Client) decodeEncryptResponse(resp Response, v interface{}) error {
	if resp.Errcode != 0 {
		return fmt.Errorf("errcode: %d, errmsg: %s", resp.Errcode, resp.Errmsg)
	}

	encrypt, err := c.encryptor.Decrypt(resp.Encrypt)
	if err != nil {
		return err
	}

	return json.NewDecoder(bytes.NewBufferString(encrypt)).Decode(v)
}

func (c *Client) decodeString(body io.Reader, output *string) error {
	b, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	*output = string(b)
	return nil
}
