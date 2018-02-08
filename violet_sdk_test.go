package violetSdk

import (
	"testing"
)

func Test_Division_1(t *testing.T) {
	res, err := GetToken("abdce")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(res)
	}
}
