package youdu

import (
	"encoding/json"
	"errors"
	"strconv"
)

const (
	groupCreateUrl = "/cgi/group/create"
	groupListUrl   = "/cgi/group/list"
)

type Group struct {
	config *Config
}

func NewGroup(config *Config) *Group {
	return &Group{
		config: config,
	}
}

// Create 创建一个群组
func (g *Group) Create(name string) (string, error) {
	accessToken, err := g.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return "", err
	}

	bodyJson, err := json.Marshal(map[string]interface{}{
		"name": name,
	})
	if err != nil {
		return "", err
	}

	encrypt, err := g.config.GetEncryptor().Encrypt(string(bodyJson))
	if err != nil {
		return "", err
	}

	resp, err := g.config.GetHttp().Post(groupCreateUrl+"?accessToken="+accessToken, map[string]interface{}{
		"appId":   g.config.AppId,
		"buin":    g.config.Buin,
		"encrypt": encrypt,
	})

	if err != nil {
		return "", err
	}

	if !resp.IsSuccess() {
		return "", errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return "", err
	}

	decrypt, err := g.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return "", err
	}

	var v map[string]string
	if err := decrypt.Unmarshal(&v); err != nil {
		return "", err
	}

	return v["id"], nil
}

func (g *Group) Delete(id int) {

}

func (g *Group) Update(id int, name string) {

}

func (g *Group) Info(id int) {

}

type GroupInfo struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Version      int    `json:"version"`
	IsDeptGroup  bool   `json:"isDeptGroup"`
	BelongDeptId int    `json:"belongDeptId"`
}

func (g *Group) List(userId ...string) ([]GroupInfo, error) {
	accessToken, err := g.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	params := map[string]string{}
	if len(userId) > 0 {
		params["userId"] = userId[0]
	}

	resp, err := g.config.GetHttp().Get(groupListUrl+"?accessToken="+accessToken, params)

	if err != nil {
		return nil, err
	}

	if !resp.IsSuccess() {
		return nil, errors.New("Response status code is " + strconv.Itoa(resp.StatusCode()))
	}

	jsonRet, err := resp.Json()
	if err != nil {
		return nil, err
	}

	decrypt, err := g.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	var v map[string][]GroupInfo
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v["groupList"], nil
}

func (g *Group) AddMember(id int, userId int) {

}

func (g *Group) RemoveMember(id int, userId int) {

}

func (g *Group) IsMember(id int, userId int) {

}
