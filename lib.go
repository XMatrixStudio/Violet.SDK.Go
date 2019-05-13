package violet

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"io"
	"strconv"
	"time"
)

// AesEncrypt 加密
func (v *Violet)aesEncrypt(text string) (string, error) {
	key := sha256.Sum256([]byte(v.Config.ClientKey))
	plaintext := []byte(text)
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", ErrorInvalidKey
	}
	cipherText := make([]byte, aes.BlockSize+len(plaintext))
	iv := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", ErrorUnknown
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plaintext)
	return hex.EncodeToString(cipherText), nil
}

// AesDecrypt 解密
func(v *Violet) aesDecrypt(cryptoText string) (string, error) {
	key :=sha256.Sum256([]byte(v.Config.ClientKey))
	cipherText, err := hex.DecodeString(cryptoText)
	if err != nil || len(cipherText) < aes.BlockSize  {
		return "", ErrorUnableDecrypt
	}
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", ErrorInvalidKey
	}
	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}

// MakeState 生成State
func (v *Violet) makeState() (string, error) {
	return v.aesEncrypt(strconv.FormatInt(time.Now().Unix(), 10))
}

// getClientSecret 获取站点密钥
func (v *Violet) getClientSecret() (string, error) {
	secret, err := v.aesEncrypt(getNowTime())
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%v&%v&%v", v.Config.ClientID, secret, getHash(v.Config.ClientID+secret+v.Config.ClientKey)), nil
}

// GetNowTime 获取时间戳
func getNowTime() string {
	return strconv.FormatInt(time.Now().Unix()*1000, 10)
}

// GetHash 获取Hash512
func getHash(str string) string {
	h := sha512.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func parsingRes(resp *resty.Response, res interface{}, errorList... error) (err error) {
	if resp.StatusCode() == 400 {
		var errRes ErrorRes
		if err = json.Unmarshal([]byte(resp.String()), &errRes); err != nil {
			return ErrorUnknown
		}
		for _, e := range errorList {
			if e.Error() == errRes.Error {
				return e
			}
		}
		return ErrorUnknown
	}
	if resp.StatusCode() / 100 != 2 {
		return ErrorServer
	} else if err = json.Unmarshal([]byte(resp.String()), &res); err != nil {
		return ErrorUnknown
	}
	return nil
}