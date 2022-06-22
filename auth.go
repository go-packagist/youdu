package youdu

import (
	"encoding/json"
	"errors"
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

type IdentifyResp struct {
	Buin   int `json:"buin"`
	Status struct {
		Code      int    `json:"code"`
		Message   string `json:"message"`
		CreatedAt string `json:"createdAt"`
	} `json:"status"`
	UserInfo struct {
		Gid        int    `json:"gid"`
		Account    string `json:"account"`
		ChsName    string `json:"chsName"`
		EngName    string `json:"engName"`
		Gender     int    `json:"gender"`
		OrgId      int    `json:"orgId"`
		Mobile     string `json:"mobile"`
		Phone      string `json:"phone"`
		Email      string `json:"email"`
		CustomAttr string `json:"customAttr"`
	} `json:"userInfo"`
}

func (a *Auth) Identify(token string) (i IdentifyResp, err error) {
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

	if err = json.Unmarshal(resp.Body(), &i); err != nil {
		return
	}

	return
}
