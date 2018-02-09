package violetSdk

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func Test_Division_1(t *testing.T) {
	res, err := GetToken("abdce")
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(res)
	}
}

func Test_Division_2(t *testing.T) {
	se, err := AesEncrypt("aes-20170416-30-1000")
	fmt.Println(se, err)
	se = se + "2333"
	sd, err := AesDecrypt(se)
	fmt.Println(sd, err)
	tm, err := strconv.ParseInt(sd, 10, 64)
	fmt.Println(tm, err)
	tms := time.Unix(tm, 0)
	fmt.Println(tms)
}
