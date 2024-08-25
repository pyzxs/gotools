package CryptUtil

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
)

// Des3Encrypt 3DES加密
func Des3Encrypt(text string, key []byte) (string, error) {
	plaintext := pad3des([]byte(text))

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	// 返回iv大小
	iv := key[:des.BlockSize]
	ciphertext := make([]byte, des.BlockSize+len(plaintext))
	mode := cipher.NewCBCEncrypter(block, iv)

	mode.CryptBlocks(ciphertext[des.BlockSize:], plaintext)

	return hex.EncodeToString(ciphertext), nil
}

// Des3Decrypt 3DES解密
func Des3Decrypt(ciphertext string, key []byte) (string, error) {
	cipherBytes, err := hex.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return "", err
	}

	// 返回iv大小
	iv := key[:des.BlockSize]
	plaintext := make([]byte, len(cipherBytes)-des.BlockSize)

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(plaintext, cipherBytes[des.BlockSize:])

	return string(unpad3des(plaintext)), nil
}

func pad3des(buf []byte) []byte {
	// 添加PKCS#5填充
	padding := des.BlockSize - (len(buf) % des.BlockSize)
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(buf, padtext...)
}

func unpad3des(buf []byte) []byte {
	// 移除PKCS#5填充
	length := len(buf)
	unpadding := int(buf[length-1])
	return buf[:length-unpadding]
}
