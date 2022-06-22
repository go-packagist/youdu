package youdu

import (
	"errors"
	"fmt"
	"strconv"
)

const (
	userGetUrl = "/cgi/user/get"
)

type User struct {
	config *Config
}

func NewUser(config *Config) *User {
	return &User{
		config: config,
	}
}

type UserInfo struct {
	Gid        int    `json:"gid"`
	UserId     string `json:"userId"`
	Name       string `json:"name"`
	Gender     int    `json:"gender"` // 性别。0表示男性，1表示女性
	Mobile     string `json:"mobile"`
	Phone      string `json:"phone"`
	Email      string `json:"email"`
	Dept       []int  `json:"dept"`
	DeptDetail []struct {
		DeptId   int    `json:"deptId"`
		DeptName string `json:"deptName"`
		Position string `json:"position"`
		Weight   int    `json:"weight"`
		SortId   int    `json:"sortId"`
	} `json:"deptDetail"`
	Attrs []interface{} `json:"attrs"`
}

// Get 获取用户信息
// see: https://youdu.im/doc/api/c01_00013.html#_6
func (u *User) Get(userId string) (*UserInfo, error) {
	accessToken, err := u.config.GetAccessTokenProvider().GetAccessToken()
	if err != nil {
		return nil, err
	}

	resp, err := u.config.GetHttp().Get(userGetUrl, map[string]string{
		"userId":      userId,
		"accessToken": accessToken,
	})

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

	decrypt, err := u.config.GetEncryptor().Decrypt(jsonRet["encrypt"].(string))
	if err != nil {
		return nil, err
	}

	fmt.Println(decrypt)

	var v *UserInfo
	if err := decrypt.Unmarshal(&v); err != nil {
		return nil, err
	}

	return v, nil
}
