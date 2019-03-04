package violetSdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty"
)

// LoginRes 登陆返回值
type LoginRes struct {
	Valid bool
	Email string
	Code  string
}

// Login 直接登陆
func (v *Violet) Login(userName, userPass string) (res LoginRes, err error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"userName": "` + userName + `", "userPass":"` + userPass + `", "clientSecret":"` + v.getClientSecret() + `"}`).
		Post(v.Config.ServerHost + "/api/Login")
	if err != nil {
		return
	}
	// 非正常的返回码
	if resp.StatusCode() != 200 {
		err = errors.New(resp.String())
		return
	}
	// 解析结果
	err = json.Unmarshal([]byte(resp.String()), &res)
	return
}

// Register 直接注册
func (v *Violet) Register(userName, userEmail, userPass string) error {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"name": "` + userName + `","email": "` + userEmail + `", "userPass":"` + userPass + `", "clientSecret":"` + v.getClientSecret() + `"}`).
		Post(v.Config.ServerHost + "/api/Register")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(resp.String())
	}
	return nil
}

// ChangePassword 直接更改密码
func (v *Violet) ChangePassword(userEmail, userPass, vCode string) error {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"vCode": "` + vCode + `","email": "` + userEmail + `", "password":"` + userPass + `", "clientSecret":"` + v.getClientSecret() + `"}`).
		Post(v.Config.ServerHost + "/api/ChangePassword")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(resp.String())
	}
	return nil
}

// GetEmailCode 获取邮箱验证码
func (v *Violet) GetEmailCode(userEmail string) error {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"email": "` + userEmail + `", "clientSecret":"` + v.getClientSecret() + `"}`).
		Post(v.Config.ServerHost + "/api/GetEmailCode")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(resp.String())
	}
	return err
}

// ValidEmail 验证邮箱
func (v *Violet) ValidEmail(userEmail, vCode string) error {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"email": "` + userEmail + `","vCode": "` + vCode + `", "clientSecret":"` + v.getClientSecret() + `"}`).
		Post(v.Config.ServerHost + "/api/ValidEmail")
	if err != nil {
		return err
	}
	if resp.StatusCode() != 200 {
		return errors.New(resp.String())
	}
	return err
}

// ---- 开放API -----

// TokenRes Token结果
type TokenRes struct {
	UserID string
	Token  string
}

// GetToken 获取Token
func (v *Violet) GetToken(code string) (res TokenRes, err error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"grantType":"authorization_code", "clientSecret":"` + v.getClientSecret() + `", "code":"` + code + `"}`).
		Post(v.Config.ServerHost + "/api/Token")
	if err != nil {
		return
	}
	if resp.StatusCode() != 200 {
		err = errors.New(resp.String())
		return
	}
	var tokenRes TokenRes
	err = json.Unmarshal([]byte(resp.String()), &tokenRes)
	return tokenRes, err
}

// UserInfoRes 用户基本信息
type UserInfoRes struct {
	Email string
	Name  string
	Info  UserInfo
}

// UserInfo 用户个性信息
type UserInfo struct {
	PublicEmail string
	Email       []string
	Bio         string
	URL         string
	Phone       string
	BirthDate   string
	Location    string
	Avatar      string
	Gender      int
}

// GetUserBaseInfo 获取用户基本信息
func (v *Violet) GetUserBaseInfo(userID, userAuth string) (res UserInfoRes, err error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"accessToken":"` + userAuth + `", "clientSecret":"` + v.getClientSecret() + `", "userId":"` + userID + `"}`).
		Post(v.Config.ServerHost + "/api/BaseData")
	if err != nil {
		return
	}
	if resp.StatusCode() != 200 {
		err = errors.New(resp.String())
		return
	}
	err = json.Unmarshal([]byte(resp.String()), &res)
	return
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

// MakeState 以当前时间生成State
func (v *Violet) MakeState() (string, error) {
	return v.AesEncrypt(strconv.FormatInt(time.Now().Unix(), 64))
}

// CheckState 检测State的正确性, 10分钟有效期
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
	if sec > 10*60 || sec < 0 {
		return false
	}
	return true
}
