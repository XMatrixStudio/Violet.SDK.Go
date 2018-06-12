package violetSdk

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty"
)

// Login 直接登陆
func (v *Violet) Login(userName, userPass string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"userName": "` + userName + `", "userPass":"` + userPass + `", "clientSecret":"` + v.getClientSecret() + `"}`).
		Post(v.Config.ServerHost + "/api/Login")
	return resp, err
}

// Register 直接注册
func (v *Violet) Register(userName, userEmail, userPass string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"name": "` + userName + `","email": "` + userEmail + `" "userPass":"` + userPass + `", "clientSecret":"` + v.getClientSecret() + `"}`).
		Post(v.Config.ServerHost + "/api/Register")
	return resp, err
}

// ChangePassword 直接更改密码
func (v *Violet) ChangePassword(userEmail, userPass, vCode string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"vCode": "` + vCode + `","email": "` + userEmail + `", "password":"` + userPass + `", "clientSecret":"` + v.getClientSecret() + `"}`).
		Post(v.Config.ServerHost + "/api/ChangePassword")
	return resp, err
}

// GetEmailCode 获取邮箱验证码
func (v *Violet) GetEmailCode(userEmail string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"email": "` + userEmail + `", "clientSecret":"` + v.getClientSecret() + `"}`).
		Post(v.Config.ServerHost + "/api/ChangePassword")
	return resp, err
}

// GetToken 获取Token
func (v *Violet) GetToken(code string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"grantType":"authorization_code", "clientSecret":"` + v.getClientSecret() + `", "code":"` + code + `"}`).
		Post(v.Config.ServerHost + "/api/Token")
	return resp, err
}

// GetUserBaseInfo 获取用户基本信息
func (v *Violet) GetUserBaseInfo(userID, userAuth string) (*resty.Response, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"accessToken":"` + userAuth + `", "clientSecret":"` + v.getClientSecret() + `", "userId":"` + userID + `"}`).
		Post(v.Config.ServerHost + "/api/BaseData")
	return resp, err
}

// GetLoginURL 获取登陆地址
func (v *Violet) GetLoginURL(redirectURL string) (url, state string) {
	state, _ = v.AesEncrypt(GetNowTime())
	url = fmt.Sprintf("%v?responseType=code&clientId=%v&state=%v&redirectUrl=%v", v.Config.LoginURL, v.Config.ClientID, state, redirectURL)
	return
}

// getClientSecret 获取站点密钥
func (v *Violet) getClientSecret() string {
	secret, _ := v.AesEncrypt(GetNowTime())
	return fmt.Sprintf("%v&%v&%v", v.Config.ClientID, secret, GetHash(secret+v.Config.ClientKey))
}

// CheckState 检测State的正确性
func (v *Violet) CheckState(state string) bool {
	b, err := v.AesDecrypt(state)
	if err != nil {
		return false
	}
	tm, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		return false
	}
	sec := time.Now().Sub(time.Unix(tm/1000, 0)).Seconds()
	if sec > 60*60 || sec < 0 {
		return false
	}
	return true
}
