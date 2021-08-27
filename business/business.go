package business

import (
	"encoding/json"
	"errors"
	"github.com/hardykay/enterprise-wechat/consts"
	"time"
)

// GetAccessToken 处理token
func GetAccessToken(corpId, corpSecret string) (data *AccessToken, err error) {
	var (
		b   []byte
		now = time.Now().Unix()
	)
	b, err = requestGet(consts.GETAccessTokenUrl, corpId, corpSecret)
	if err != nil {
		return
	}
	data = new(AccessToken)
	err = json.Unmarshal(b, data)
	if err != nil {
		return
	}
	if data.ErrCode != 0 {
		err = errors.New(data.ErrMsg)
	}
	// 用于计算剩余时间
	data.CreatedAt = now
	return
}

// GetUserId  获取Userid
func GetUserId(token, code string) (data *UserId, err error) {
	var (
		b []byte
	)
	b, err = requestGet(consts.GETUserIdUrlUrl, token, code)
	if err != nil {
		return
	}
	data = new(UserId)
	err = json.Unmarshal(b, data)
	if err != nil {
		return
	}
	if data.ErrCode != 0 {
		err = errors.New(data.ErrMsg)
	}
	return
}

// GetCode2Session 获取Code2Session
func GetCode2Session(token, code string) (data *Code2Session, err error) {
	var (
		b []byte
	)
	b, err = requestGet(consts.GETCode2SessionUrl, token, code)
	if err != nil {
		return
	}
	data = new(Code2Session)
	err = json.Unmarshal(b, data)
	if err != nil {
		return
	}
	if data.ErrCode != 0 {
		err = errors.New(data.ErrMsg)
	}
	return
}

// GetUserDetail 获取成员详情
func GetUserDetail(token, code string) (data *UserDetail, err error) {
	var (
		b []byte
	)
	b, err = requestGet(consts.GETUserDetailUrl, token, code)
	if err != nil {
		return
	}
	data = new(UserDetail)
	err = json.Unmarshal(b, data)
	if err != nil {
		return
	}
	if data.ErrCode != 0 {
		err = errors.New(data.ErrMsg)
	}
	return
}
