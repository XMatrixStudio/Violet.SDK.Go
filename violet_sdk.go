package violetSdk

import (
	"github.com/go-resty/resty"
)

// GetToken 获取Token
func GetToken(code string) (string, error) {
	resp, err := resty.R().
		SetHeader("Content-Type", "application/json").
		SetBody(`{"grantType":"authorization_code", "clientSecret":"testpass", "code":"` + code + `"}`).
		Get("https://oauth.xmatrix.studio/api/v2/verify/Token")
	return resp.String(), err
}
