package violet

import (
	"fmt"
	"testing"
)

var violet Violet

func Test_Division_1(t *testing.T) {
	violet = NewViolet(Config{
		ClientID:   "5ccc67b6a9eb661b34935229",
		ClientKey:  "4ew5i25ozqy2znxxmihuhjov",
		ServerHost: "http://localhost:3000/api",
		LoginURL:   "http://localhost:3000/account/auth",
	})
	url, state, err := violet.GetLoginURL(
		"https://blog.zhenly.cn/auth",
		AuthOption{Scopes: ScopeTypes{ScopeInfo}, QuickMode: false})
	fmt.Println(url, state, err)
}
