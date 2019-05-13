package violet

import (
	"testing"
)

var violet = NewViolet(Config{
	ClientID:   "5cd91e09131df3397c35b079",
	ClientKey:  "d9d91xret2h75hrwfscjcahq",
	ServerHost: "http://localhost:3000/api",
	LoginURL:   "http://localhost:3000/account/auth",
})

var code = "a18d8b33df24783d736a3eb7e234923ef336e0bbad3873205499c4801de495e6983a10c800f7df3c1e843e5b59f72183a77686bd53adb4a788dd1b26827daeb3a9bf2108b7377bc017f9e5eddd26c9d4d3878cf7efc7bf14496976901f4dc1711f"

func getViolet() *Violet {
	return violet
}

func TestViolet_GetLoginURL(t *testing.T) {
	url, _, err := getViolet().GetLoginURL(
		"https://blog.zhenly.cn/auth",
		AuthOption{Scopes: ScopeTypes{ScopeInfo}, QuickMode: false})
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
}
