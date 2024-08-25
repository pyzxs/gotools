package CryptUtil

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

var (
	DEFAULT_AES_KEY = "1234567890123456"
)

// AesEcbEncrypt Ecb模式加密 电码本模式（Electronic Codebook Book (ECB)）
func AesEcbEncrypt(src []byte, key []byte) (encrypted []byte) {
	cp, _ := aes.NewCipher(GenerateKey(key))
	length := (len(src) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, src)
	pad := byte(len(plain) - len(src))
	for i := len(src); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted = make([]byte, len(plain))
	// 分组分块加密
	for bs, be := 0, cp.BlockSize(); bs <= len(src); bs, be = bs+cp.BlockSize(), be+cp.BlockSize() {
		cp.Encrypt(encrypted[bs:be], plain[bs:be])
	}

	return encrypted
}

// AesEcbDecrypt Ecb模式解密
func AesEcbDecrypt(encrypted []byte, key []byte) (decrypted []byte) {
	cp, _ := aes.NewCipher(GenerateKey(key))
	decrypted = make([]byte, len(encrypted))
	//
	for bs, be := 0, cp.BlockSize(); bs < len(encrypted); bs, be = bs+cp.BlockSize(), be+cp.BlockSize() {
		cp.Decrypt(decrypted[bs:be], encrypted[bs:be])
	}

	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}

	return decrypted[:trim]
}

// AesCbcEncrypt  CBC模式加密  密码分组链接模式
func AesCbcEncrypt(origData, key []byte) []byte {
	// 分组秘钥
	// NewCipher该函数限制了输入k的长度必须为16, 24或者32
	block, _ := aes.NewCipher(key)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 补全码
	origData = pkcs7Padding(origData, blockSize)
	// 加密模式
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	// 创建数组
	cryted := make([]byte, len(origData))
	// 加密
	blockMode.CryptBlocks(cryted, origData)

	return cryted
}

// AesCbcDecrypt CBC模式解密
func AesCbcDecrypt(crytedByte, key []byte) []byte {
	// 转成字节数组
	k := []byte(key)
	// 分组秘钥
	block, _ := aes.NewCipher(k)
	// 获取秘钥块的长度
	blockSize := block.BlockSize()
	// 加密模式
	blockMode := cipher.NewCBCDecrypter(block, k[:blockSize])
	// 创建数组
	orig := make([]byte, len(crytedByte))
	// 解密
	blockMode.CryptBlocks(orig, crytedByte)
	// 去补全码
	return pkcs7UnPadding(orig)
}

// AesCtrCrypt 计算器模式（Counter (CTR)） 加解密同一函数
func AesCtrCrypt(plainText []byte, key []byte) ([]byte, error) {

	//1. 创建cipher.Block接口
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//2. 创建分组模式，在crypto/cipher包中
	iv := bytes.Repeat([]byte("1"), block.BlockSize())
	stream := cipher.NewCTR(block, iv)
	//3. 加密
	dst := make([]byte, len(plainText))
	stream.XORKeyStream(dst, plainText)

	return dst, nil
}

// AesCfbEncrypt 密码反馈模式（Cipher FeedBack (CFB)）
func AesCfbEncrypt(origData []byte, key []byte) (encrypted []byte) {
	block, err := aes.NewCipher(key)
	if err != nil {
		//panic(err)
	}
	encrypted = make([]byte, aes.BlockSize+len(origData))
	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		//panic(err)
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], origData)
	return encrypted
}

// AesCfbDecrypt CFB模式解密
func AesCfbDecrypt(encrypted []byte, key []byte) (decrypted []byte) {
	block, _ := aes.NewCipher(key)
	if len(encrypted) < aes.BlockSize {
		panic("ciphertext too short")
	}
	iv := encrypted[:aes.BlockSize]
	encrypted = encrypted[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(encrypted, encrypted)
	return encrypted
}

// AesOfbEncrypt 输出反馈模式（Output FeedBack (OFB)）
func AesOfbEncrypt(data []byte, key []byte) ([]byte, error) {
	data = pkcs7Padding(data, aes.BlockSize)
	block, _ := aes.NewCipher([]byte(key))
	out := make([]byte, aes.BlockSize+len(data))
	iv := out[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewOFB(block, iv)
	stream.XORKeyStream(out[aes.BlockSize:], data)
	return out, nil
}

// AesOfbDecrypt OFB模式解密
func AesOfbDecrypt(data []byte, key []byte) ([]byte, error) {
	block, _ := aes.NewCipher([]byte(key))
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	if len(data)%aes.BlockSize != 0 {
		return nil, fmt.Errorf("data is not a multiple of the block size")
	}

	out := make([]byte, len(data))
	mode := cipher.NewOFB(block, iv)
	mode.XORKeyStream(out, data)

	out = pkcs7UnPadding(out)
	return out, nil
}

// GenerateKey 生成key
func GenerateKey(key []byte) (genKey []byte) {
	genKey = make([]byte, 16)
	copy(genKey, key)
	for i := 16; i < len(key); {
		for j := 0; j < 16 && i < len(key); j, i = j+1, i+1 {
			genKey[j] ^= key[i]
		}
	}
	return genKey
}

// pkcs7Padding 补码
// AES加密数据块分组长度必须为128bit(byte[16])，
// 密钥长度可以是128bit(byte[16])、192bit(byte[24])、256bit(byte[32])中的任意一个。
func pkcs7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// 去码
func pkcs7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
