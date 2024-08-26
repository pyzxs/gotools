package CryptUtil

import (
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
)

// Sha1 计算字符串的Sha1哈希值
func Sha1(str string) []byte {
	h := sha1.New()
	_, _ = h.Write([]byte(str))
	return h.Sum(nil)
}

// HmacSha1 使用HMAC方法计算字符串的sha1哈希值
func HmacSha1(key string, data string) []byte {
	mac := hmac.New(sha1.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

// HmacSha1ToString 使用HMAC方法计算字符串的sha1哈希，输出小写十六进制
func HmacSha1ToString(key string, data string) string {
	return hex.EncodeToString(HmacSha1(key, data))
}

// Sha256 计算字符串的Sha256哈希值
func Sha256(str string) []byte {
	h := sha256.New()
	_, _ = h.Write([]byte(str))
	return h.Sum(nil)
}

// HmacSha256 使用HMAC方法计算字符串的sha256哈希值
func HmacSha256(key string, data string) []byte {
	mac := hmac.New(sha256.New, []byte(key))
	_, _ = mac.Write([]byte(data))

	return mac.Sum(nil)
}

// HmacSha256ToString 使用HMAC方法计算字符串的sha256哈希，输出小写十六进制
func HmacSha256ToString(key string, data string) string {
	return hex.EncodeToString(HmacSha256(key, data))
}
