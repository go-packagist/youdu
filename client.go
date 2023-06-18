package youdu

import (
	"net/http"
)

type Client struct {
	config     *Config
	httpClient *http.Client
	encryptor  *Encryptor
	token      *Token
}

func NewClient(config *Config) *Client {
	return &Client{
		config:     config,
		httpClient: &http.Client{},
		encryptor:  NewEncryptor(config.AppId, config.AesKey),
	}
}
