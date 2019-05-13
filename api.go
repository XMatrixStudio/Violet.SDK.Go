package violet

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	resty "gopkg.in/resty.v1"
)

// Login 直接登陆
func (v *Violet) Login(userName, userPass string) (res LoginRes, err error) {
	return res, errors.New("not_implemented")
	//secret, err := v.getClientSecret()
	//if err != nil {
	//	return
	//}
	//resp, err := resty.R().
	//	SetHeader("Content-Type", "application/json").
	//	SetBody(`{"userName": "` + userName + `", "userPass":"` + userPass + `", "clientSecret":"` + secret + `"}`).
	//	Post(v.Config.ServerHost + "/api/Login")
	//if err != nil {
	//	return
	//}
	//// 非正常的返回码
	//if resp.StatusCode() != 200 {
	//	err = errors.New(resp.String())
	//	return
	//}
	//// 解析结果
	//err = json.Unmarshal([]byte(resp.String()), &res)
	//return
}

// Register 直接注册
func (v *Violet) Register(userName, userEmail, userPass string) error {
	return errors.New("not_implemented")
	//secret, err := v.getClientSecret()
	//if err != nil {
	//	return err
	//}
	//resp, err := resty.R().
	//	SetHeader("Content-Type", "application/json").
	//	SetBody(`{"name": "` + userName + `","email": "` + userEmail +
	//		`", "userPass":"` + userPass + `", "clientSecret":"` + secret + `"}`).
	//	Post(v.Config.ServerHost + "/api/Register")
	//if err != nil {
	//	return err
	//}
	//if resp.StatusCode() != 200 {
	//	return errors.New(resp.String())
	//}
	//return nil
}

// ChangePassword 直接更改密码
func (v *Violet) ChangePassword(userEmail, userPass, vCode string) error {
	return errors.New("not_implemented")
	//secret, err := v.getClientSecret()
	//if err != nil {
	//	return err
	//}
	//resp, err := resty.R().
	//	SetHeader("Content-Type", "application/json").
	//	SetBody(`{"vCode": "` + vCode + `","email": "` + userEmail +
	//		`", "password":"` + userPass + `", "clientSecret":"` + secret + `"}`).
	//	Post(v.Config.ServerHost + "/api/ChangePassword")
	//if err != nil {
	//	return err
	//}
	//if resp.StatusCode() != 200 {
	//	return errors.New(resp.String())
	//}
	//return nil
}

// GetEmailCode 获取邮箱验证码
func (v *Violet) GetEmailCode(userEmail string) error {
	return errors.New("not_implemented")
	//secret, err := v.getClientSecret()
	//if err != nil {
	//	return err
	//}
	//resp, err := resty.R().
	//	SetHeader("Content-Type", "application/json").
	//	SetBody(`{"email": "` + userEmail + `", "clientSecret":"` + secret + `"}`).
	//	Post(v.Config.ServerHost + "/api/GetEmailCode")
	//if err != nil {
	//	return err
	//}
	//if resp.StatusCode() != 200 {
	//	return errors.New(resp.String())
	//}
	//return err
}

// ValidEmail 验证邮箱
func (v *Violet) ValidEmail(userEmail, vCode string) error {
	return errors.New("not_implemented")
	//secret, err := v.getClientSecret()
	//if err != nil {
	//	return err
	//}
	//resp, err := resty.R().
	//	SetHeader("Content-Type", "application/json").
	//	SetBody(`{"email": "` + userEmail + `","vCode": "` + vCode + `", "clientSecret":"` + secret + `"}`).
	//	Post(v.Config.ServerHost + "/api/ValidEmail")
	//if err != nil {
	//	return err
	//}
	//if resp.StatusCode() != 200 {
	//	return errors.New(resp.String())
	//}
	//return err
}

// ---- 开放API -----

/*
GetToken 获取Token

@param code 	授权成功后返回的Code

@return res 	用户名以及对应的Token
@return err 	错误信息
				ErrorInvalidKey
				ErrorNetwork
				ErrorTimeOutCode
				ErrorInvalidCode
				ErrorUnknown
*/
func (v *Violet) GetToken(code string) (res TokenRes, err error) {
	secret, err := v.getClientSecret()
	if err != nil {
		return
	}
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"grantType":"authorization_code", "appSecret":"` + secret + `", "code":"` + code + `"}`).
		Post(v.Config.ServerHost + "/verify/token")
	if err != nil {
		return res, ErrorNetwork
	}
	err = parsingRes(resp, &res, ErrorTimeoutCode, ErrorInvalidCode)
	return
}

/*
GetUserInfo 获取用户信息

@param token 	该用户的 token

@return res 	用户信息
@return err 	错误信息
				ErrorInvalidKey
				ErrorNetwork
				ErrorTimeoutToken
				ErrorInvalidToken
				ErrorUnknown
*/
func (v *Violet) GetUserInfo(token string) (res UserInfoRes, err error) {
	secret, err := v.getClientSecret()
	if err != nil {
		return
	}
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetQueryParam("token", token).
		SetQueryParam("appSecret", secret).
		Get(v.Config.ServerHost + "/users/info")
	if err != nil {
		return res, ErrorNetwork
	}
	err = parsingRes(resp, &res, ErrorInvalidToken, ErrorTimeoutToken)
	return
}

/*
GetLoginURL 获取登陆地址

@param redirectURL 	回调地址
@param options 		可选选项

@return url 		用户需要跳转登陆的地址
@return state 		状态, 须与当前 Session 绑定，授权完成后须验证其有效性，防止 CSRF 攻击
@return err 		错误
					ErrorInvalidKey
					ErrorUnknown
*/
func (v *Violet) GetLoginURL(redirectURL string, options ...AuthOption) (authURL, state string, err error) {
	scopes := ScopeTypes{scopeBase}
	quickMode := true
	if len(options) != 0 {
		scopes = append(scopes, options[0].Scopes...)
		quickMode = options[0].QuickMode
	}
	state, err = v.makeState()
	authURL = fmt.Sprintf("%v/verify/authorize?responseType=code&appId=%v&state=%v&redirectUrl=%v&quickMode=%v&scope=%v",
		v.Config.ServerHost, v.Config.ClientID, state,
		url.QueryEscape(redirectURL), quickMode, url.QueryEscape(strings.Join(scopes.String(), ",")))
	return
}

/*
CheckState 		检测State的正确性, 10分钟有效期

@param	state 	需要检测的 state

@return	error	错误
				ErrorInvalidKey
				ErrorUnknown
				ErrorInvalidState
				ErrorTimeoutState
*/
func (v *Violet) CheckState(state string) error {
	b, err := v.aesDecrypt(state)
	if err != nil {
		return err
	}
	tm, err := strconv.ParseInt(b, 10, 64)
	if err != nil {
		return ErrorInvalidState
	}
	sec := time.Now().Sub(time.Unix(tm, 0)).Seconds()
	if sec > 10*60 {
		return ErrorTimeoutState
	}
	return nil
}
