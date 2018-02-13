package violetSdk

import (
	"fmt"
	"testing"
)

var globeState = ""

func Test_Division_1(t *testing.T) {
	fmt.Println("GetToken: ")
	res, err := GetToken("code")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func Test_Division_2(t *testing.T) {
	fmt.Println("GetUserBaseInfo: ")
	res, err := GetUserBaseInfo("userID", "userAuth")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func Test_Division_3(t *testing.T) {
	fmt.Println("ClientSecret: ", getClientSecret())
}

func Test_Division_4(t *testing.T) {
	url, state := GetLoginURL("redirectURL")
	fmt.Println("GetLoginURL: ", url)
	fmt.Println("GetLoginState: ", state)
	globeState = state
}

func Test_Division_5(t *testing.T) {
	fmt.Println("CheckState[true]: ", CheckState(globeState))
	fmt.Println("CheckState[false]: ", CheckState("state"))
}
