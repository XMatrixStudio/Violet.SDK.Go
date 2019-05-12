package violet

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"errors"
	"io"
	"strconv"
	"time"
)

// AesEncrypt 加密
func (v *Violet)AesEncrypt(text string) (string, error) {
	key := sha256.Sum256([]byte(v.Config.ClientKey))
	plaintext := []byte(text)
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)
	return hex.EncodeToString(ciphertext), nil
}

// AesDecrypt 解密
func(v *Violet) AesDecrypt(cryptoText string) (string, error) {
	key :=sha256.Sum256([]byte(v.Config.ClientKey))
	ciphertext, err := hex.DecodeString(cryptoText)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key[:])
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)
	return string(ciphertext), nil
}

// GetNowTime 获取时间戳
func GetNowTime() string {
	return strconv.FormatInt(time.Now().Unix()*1000, 10)
}

// GetHash 获取Hash512
func GetHash(str string) string {
	h := sha512.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
