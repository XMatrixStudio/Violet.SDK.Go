package violetSdk

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty"
)

// Login 直接登陆
func Login(userName, userPass string) (string, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"userName": "` + userName + `", "userPass":"` + userPass + `", "clientSecret":"` + getClientSecret() + `"}`).
		Post(ServerHost + "/api/Login")
	return resp.String(), err
}

// Register 直接注册
func Register(userName, userEmail, userPass string) (string, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"name": "` + userName + `","email": "` + userEmail + `" "userPass":"` + userPass + `", "clientSecret":"` + getClientSecret() + `"}`).
		Post(ServerHost + "/api/Register")
	return resp.String(), err
}

// ChangePassword 直接更改密码
func ChangePassword(userEmail, userPass, vCode string) (string, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"vCode": "` + vCode + `","email": "` + userEmail + `", "password":"` + userPass + `", "clientSecret":"` + getClientSecret() + `"}`).
		Post(ServerHost + "/api/ChangePassword")
	return resp.String(), err
}

// GetEmailCode 获取邮箱验证码
func GetEmailCode(userEmail string) (string, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"email": "` + userEmail + `", "clientSecret":"` + getClientSecret() + `"}`).
		Post(ServerHost + "/api/ChangePassword")
	return resp.String(), err
}

// GetToken 获取Token
func GetToken(code string) (string, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"grantType":"authorization_code", "clientSecret":"` + getClientSecret() + `", "code":"` + code + `"}`).
		Get(ServerHost + "/verify/Token")
	return resp.String(), err
}

// GetUserBaseInfo 获取用户基本信息
func GetUserBaseInfo(userID, userAuth string) (string, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"accessToken":"` + userAuth + `", "clientSecret":"` + getClientSecret() + `", "userId":"` + userID + `"}`).
		Get(ServerHost + "/users/BaseData")
	return resp.String(), err
}

// GetLoginURL 获取登陆地址
func GetLoginURL(redirectURL string) (url, state string) {
	state, _ = AesEncrypt(GetNowTime())
	url = fmt.Sprintf("%v?responseType=code&clientId=%v&state=%v&redirectUrl=%v", LoginURL, ClientID, state, redirectURL)
	return
}

// getClientSecret 获取站点密钥
func getClientSecret() string {
	secret, _ := AesEncrypt(GetNowTime())
	return fmt.Sprintf("%v&%v&%v", ClientID, secret, GetHash(secret+string(ClientKey[:24])))
}

// CheckState 检测State的正确性
func CheckState(state string) bool {
	b, err := AesDecrypt(state)
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
