package youdu

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	identifyUrl = "/cgi/identify"
)

type Auth struct {
	config *Config
}

func NewAuth(config *Config) *Auth {
	return &Auth{
		config: config,
	}
}

func (a *Auth) Identify(token string) (i interface{}, err error) {
	resp, err := a.config.GetHttp().Get(identifyUrl, map[string]string{
		"token": token,
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

	fmt.Println(jsonRet)

	if jsonRet["errcode"].(float64) != 0 {
		err = errors.New(jsonRet["errmsg"].(string))
		return
	}

	return

	// decrypt, err := m.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	// if err != nil {
	// 	return
	// }
	//
	// var v MediaInfo
	// if err = decrypt.Unmarshal(&v); err != nil {
	// 	return
	// }
	//
	// return v, nil
}
