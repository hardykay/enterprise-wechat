package consts

const (
	GETAccessTokenUrl  = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
	GETUserIdUrlUrl    = "https://qyapi.weixin.qq.com/cgi-bin/user/getuserinfo?access_token=%s&code=%s"
	GETCode2SessionUrl = "https://qyapi.weixin.qq.com/cgi-bin/miniprogram/jscode2session?access_token=%s&js_code=%s&grant_type=authorization_code"
	GETUserDetailUrl   = "https://qyapi.weixin.qq.com/cgi-bin/user/get?access_token=%s&userid=%s"
)
