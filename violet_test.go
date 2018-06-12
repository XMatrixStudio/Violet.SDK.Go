package violetSdk

import (
	"fmt"
	"testing"
)

var globeState = ""
var violet Violet

func Test_Division_1(t *testing.T) {
	violet = NewViolet(Config{
		ClientID:   "xxxxx",
		ClientKey:  "xxxxx",
		ServerHost: "https://oauth.xmatrix.studio/api/v2",
		LoginURL:   "https://oauth.xmatrix.studio/Verify/Authorize",
	})
	s := violet.getClientSecret()
	fmt.Println(s)
}
