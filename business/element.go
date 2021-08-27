package business

// AccessToken 企业微信token
type AccessToken struct {
	ErrCode     int    `json:"errcode"` //
	ErrMsg      string `json:"errmsg"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
	CreatedAt   int64  `json:"created_at"`
}

// UserId 企业微信userid
type UserId struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	UserId  string `json:"UserId"`
	OpenId  string `json:"OpenId"`
}

// Code2Session 临时登录凭证校验接口是一个服务端HTTPS 接口，开发者服务器使用临时登录凭证code获取 session_key、用户userid以及用户所在企业的corpid等信息。
type Code2Session struct {
	ErrCode    int    `json:"errcode"`
	ErrMsg     string `json:"errmsg"`
	UserId     string `json:"UserId"`
	CorpId     string `json:"corpid"`
	SessionKey string `json:"session_key"`
}

// UserDetail 企业微信用户详细信息
type UserDetail struct {
	ErrCode          int    `json:"errcode"`
	ErrMsg           string `json:"errmsg"`
	UserId           string `json:"UserId"`
	Name             string `json:"name"`
	Department       []int  `json:"department"`
	Order            []int  `json:"order"`
	Position         string `json:"position"`
	Mobile           string `json:"mobile"`
	Gender           string `json:"gender"`
	Email            string `json:"email"`
	Avatar           string `json:"avatar"`
	Alias            string `json:"alias"`
	Address          string `json:"address"`
	OpenUserid       string `json:"open_userid"`
	MainDepartment   int    `json:"main_department"`
	Status           int    `json:"status"`
	ExternalPosition string `json:"external_position"`
}
