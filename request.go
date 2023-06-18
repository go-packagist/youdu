package youdu

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

func (c *Client) newRequest(method, path string, request interface{}) (*http.Request, error) {
	body, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, c.config.Api+path, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	return req, nil
}

func (c *Client) newRequestWithToken(method, path string, request interface{}) (*http.Request, error) {
	req, err := c.newRequest(method, path, request)
	if err != nil {
		return nil, err
	}

	token, err := c.GetToken()
	if err != nil {
		return nil, err
	}

	req.URL.Query().Add("AccessToken", token.AccessToken)

	return req, nil
}

func (c *Client) sendRequest(req *http.Request, response interface{}) error {
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return c.handleErrorResp(resp)
	}

	return c.decodeResponse(resp.Body, response)
}

func (c *Client) handleErrorResp(resp *http.Response) error {
	return fmt.Errorf("http status: %s, http body: %s", resp.Status, resp.Body)
}
