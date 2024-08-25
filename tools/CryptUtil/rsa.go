package CryptUtil

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

const (
	DEFAULT_PATH         = "./"
	DEFAULT_PUBLIC_FILE  = "rsa_public_key.pem"
	DEFAULT_PRIVATE_FILE = "rsa_private_key"
)

type RsaCert struct {
	PrivateKey []byte
	PublicKey  []byte
}

func NewRsaCertFromKey(publiKey, privateKey []byte) *RsaCert {
	return &RsaCert{
		PrivateKey: privateKey,
		PublicKey:  publiKey,
	}
}

// 公钥: 根据私钥生成
// openssl rsa -in rsa_private_key.pem -pubout -out rsa_public_key.pem
func (r *RsaCert) encrypt(origData []byte) ([]byte, error) {
	//解密pem格式的公钥
	block, _ := pem.Decode(r.PublicKey)
	if block == nil {
		return nil, errors.New("public key error")
	}
	// 解析公钥
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)
	//加密
	return rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
}

// 解密
// openssl genrsa -out rsa_private_key.pem 1024
func (r *RsaCert) Decrypt(ciphertext []byte) ([]byte, error) {
	//解密
	block, _ := pem.Decode(r.PrivateKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}
	//解析PKCS1格式的私钥
	priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	return rsa.DecryptPKCS1v15(rand.Reader, priv, ciphertext)
}
