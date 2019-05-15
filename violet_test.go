package violet

import (
	"testing"
)

var violet = NewViolet(Config{
	ClientID:   "5cda7372c254be1542ae4f43",
	ClientKey:  "gu1bzgbsf2h16ahfjrpwxbv4",
	ServerHost: "https://love.zhenly.cn/api",
})

var code = "cc52b5372f3c3e9056b27c8b2a36942bcex73f89aed05d074b602d290d0d9c2fa9bad20b77eb56a6e00c4bcfe6da21f7cff19d76118e0ee8b16e9b440a86c130c83abdbab61b1da112d1a54e4782f12969ce8ed709e1458cb926a91189e6c59717"

func getViolet() *Violet {
	return violet
}

func TestViolet_GetLoginURL(t *testing.T) {
	url, _, err := getViolet().GetLoginURL(
		"http://localhost:8080/api/session/violet",
		AuthOption{Scopes: ScopeTypes{ScopeInfo}, QuickMode: true})

	if err != nil {
		t.Error(err)
	}
	t.Log("url:", url)
}

func TestViolet_CheckState(t *testing.T) {
	state, err := getViolet().makeState()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("state:", state)
	err = getViolet().CheckState(state)
	if err != nil {
		t.Error(err)
	}
}

func TestViolet_getClientSecret(t *testing.T) {
	secret, err := getViolet().getClientSecret()
	if err != nil {
		t.Error(err)
	}
	t.Log("secret:", secret)
}

func TestViolet_GetToken(t *testing.T) {
	token, err := getViolet().GetToken(code)
	if err != nil {
		t.Error(err)
	}
	t.Log("token:", token.Token)
	t.Log("userID:", token.UserID)
}

func TestViolet_GetUserInfo(t *testing.T) {
	token, err := getViolet().GetToken(code)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log("token:", token.Token)
	info, err := getViolet().GetUserInfo(token.Token)
	if err != nil {
		t.Error(err)
	}
	t.Log("id:", info.ID)
	t.Log("bio:", info.Bio)
	t.Log("nickname:", info.Nickname)
	t.Log("avatar:", info.Avatar)
	t.Log("email:", info.Email)
	t.Log("phone:", info.Phone)
	t.Log("location:", info.Location)
	t.Log("birthday:", info.Birthday)
	t.Log("url:", info.URL)
	t.Log("gender:", info.Gender)
}
