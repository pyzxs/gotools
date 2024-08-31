package CryptUtil

import (
	"crypto/aes"
	"crypto/cipher"
)

const (
	BLOCK_COUNT = "12345678abcdefgh"
)

// AEC加密和解密（CRT模式）
func AesCtrCrypt(plainText []byte, key []byte) ([]byte, error) {

	//指定加密、解密算法为AES，返回一个AES的Block接口对象
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	//指定计数器,长度必须等于block的块尺寸
	count := []byte(BLOCK_COUNT)
	//指定分组模式
	blockMode := cipher.NewCTR(block, count)
	//执行加密、解密操作
	message := make([]byte, len(plainText))
	blockMode.XORKeyStream(message, plainText)

	//返回明文或密文
	return message, nil
}
