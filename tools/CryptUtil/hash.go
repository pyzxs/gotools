package CryptUtil

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"hash"
)

// MD5 md5加密
func MD5(data string) string {
	return hashFromFunc(md5.New, data)
}

// HmacFromMD5 HMAC是密钥相关的哈希运算消息认证码
func HmacFromMD5(key, data string) string {
	return hmacFromFunc(key, data, md5.New)
}

// HmacFromSha1 通过sha1算法
func HmacFromSha1(key, data string) string {
	return hmacFromFunc(key, data, sha1.New)
}

// HmacFromSha256 通过sha256算法
func HmacFromSha256(key, data string) string {
	return hmacFromFunc(key, data, sha256.New)
}

// HmacFromSha512 通过sha512算法
func hmacFromSha512(key, data string) string {
	return hmacFromFunc(key, data, sha512.New)
}

// Sha1 sha1算列
func Sha1(data string) string {
	return hashFromFunc(sha1.New, data)
}

// Sha256 sha256散列
func Sha256(data string) string {
	return hashFromFunc(sha256.New, data)
}

// Sha512 sha512散列
func Sha512(data string) string {
	return hashFromFunc(sha512.New, data)
}

func hashFromFunc(init func() hash.Hash, data string) string {
	h := init()
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}

func hmacFromFunc(key, data string, fn func() hash.Hash) string {
	h := hmac.New(fn, []byte(key))
	h.Write([]byte(data))
	return hex.EncodeToString(h.Sum(nil))
}
