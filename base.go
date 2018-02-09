package violetSdk

import (
	"crypto/sha512"
	"fmt"

	"github.com/go-resty/resty"
)

// GetToken 获取Token
func GetToken(code string) (string, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"grantType":"authorization_code", "clientSecret":"` + getClientSecret() + `", "code":"` + code + `"}`).
		Get(ServerHost + "/verify/Token")
	return resp.String(), err
}

// GetUserBaseInfo 获取用户基本信息
func GetUserBaseInfo(userId, userAuth string) (string, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"accessToken":"` + userAuth + `", "clientSecret":"` + getClientSecret() + `", "userId":"` + userId + `"}`).
		Get(ServerHost + "/user/BaseData")
	return resp.String(), err
}

// getClientSecret 获取站点密钥
func getClientSecret() string {
	secret, _ := AesEncrypt(GetNowTime())
	h := sha512.New()
	h.Write([]byte(secret + string(ClientKey[:24])))
	bs := h.Sum(nil)
	return fmt.Sprintf("%v&%v&%v", ClientID, secret, bs)
}

// GetLoginURL 获取登陆地址
func GetLoginURL(redirectURL string) (url, state string) {
	state, _ = AesEncrypt(GetNowTime())
	url = fmt.Sprintf("%v?responseType=code&clientId=%v&state=%v&redirectUrl=%v", LoginUrl, ClientID, state, redirectURL)
	return
}

/*
func CheckState(state, string) bool {
	state, _ = AesDecrypt(state)

}

async function checkState (state) {
  try {
    state = decrypt(state)
    if (!state) throw new Error() // 错误的state
    let validTime = new Date(state)
    if (Number.isNaN(validTime.getTime())) throw new Error() // 不合法的state
    if (new Date().getTime() - validTime > 1000 * 60 * 60) throw new Error() // 过期的state
    return true
  } catch (error) {
    return false
  }
}

*/
