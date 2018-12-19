package util

import (
	"encoding/base64"
	"crypto/sha256"
	"fmt"
	"errors"
	"crypto/rand"
)

//base64
const (
	Table = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)
var coder = base64.NewEncoding(Table)
func Base64Encode(src []byte) []byte {
	return []byte(coder.EncodeToString(src))
}
func Base64Decode(src []byte) ([]byte, error) {
	return coder.DecodeString(string(src))
}



//获取盐
func saltString() (salt string, error error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", errors.New(err.Error())
	}
	return fmt.Sprintf("%x", b), nil
}

//对密码进行sha256
func EncryptPasswordWithSalt(password, salt string) (hashedPwd string, error error) {
	sha_256 := sha256.New()
	_, err := sha_256.Write([]byte(password + salt))
	if err != nil {
		return "", errors.New(err.Error())
	}
	return fmt.Sprintf("%x", sha_256.Sum(nil)), nil
}

//加密密码 返回加密结果 以及使用的盐
func EncryptPassword(password string) (hashedPwd string, salt string, error error) {
	saltStr, err := saltString()
	if err != nil {
		return "", "", errors.New("server err")
	}

	hashedPd, err := EncryptPasswordWithSalt(password, saltStr)
	if err != nil {
		return "", "", errors.New("server err")
	}
	return hashedPd, saltStr, nil
}