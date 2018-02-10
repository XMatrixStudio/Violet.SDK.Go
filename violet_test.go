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
	se, err := AesEncrypt(strconv.FormatInt(time.Now().Unix()*1000, 10))
	fmt.Println("t: ", strconv.FormatInt(time.Now().Unix(), 10))
	fmt.Println(se, err)
	valid := CheckState(se)
	fmt.Println(valid)
}

func Test_Division_3(t *testing.T) {
	fmt.Println(GetHash("hello, world"))
}
