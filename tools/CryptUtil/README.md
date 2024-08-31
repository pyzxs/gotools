# CryptUtil 项目包
- [AES](#AES)
- [DES](#DES)
- [3DES](#DES)
- [RSA](#RSA)
- [HMAC](#HMAC)
- [HASH](#HASH)
- [BCRYPT](#BCRYPT)
- [CTR](#CTR)


## Usage

### AES

秘钥长度 16/24/32 个字符 (128/192/256 bits)

AES-ECB:

```go 
src := []byte("123456")
key := []byte("1234567890123456")
dst , _ := CryptUtil.AesECBEncrypt(src, key, CryptUtil.PKCS7_PADDING)
fmt.Printf(base64.StdEncoding.EncodeToString(dst))  // yXVUkR45PFz0UfpbDB8/ew==

dst , _ = CryptUtil.AesECBDecrypt(dst, key, CryptUtil.PKCS7_PADDING)
fmt.Println(string(dst)) // 123456
```

AES-CBC:

```go
src := []byte("123456")
key := []byte("1234567890123456")
iv := []byte("1234567890123456")
dst , _ := CryptUtil.AesCBCEncrypt(src, key, iv, CryptUtil.PKCS7_PADDING)
fmt.Println(base64.StdEncoding.EncodeToString(dst)) // 1jdzWuniG6UMtoa3T6uNLA==

dst , _ = CryptUtil.AesCBCDecrypt(dst, key, iv, CryptUtil.PKCS7_PADDING)
fmt.Println(string(dst)) // 123456
```

### DES

秘钥长度8个字符 (64 bits)

DES-ECB:

```go
CryptUtil.DesECBEncrypt(src, key, CryptUtil.PKCS7_PADDING)
CryptUtil.DesECBDecrypt(src, key, CryptUtil.PKCS7_PADDING)
```

DES-CBC:

```go
CryptUtil.DesCBCEncrypt(src, key, iv, CryptUtil.PKCS7_PADDING)
CryptUtil.DesCBCDecrypt(src, key, iv, CryptUtil.PKCS7_PADDING)
```

### 3DES

秘钥长度是24个字符 (192 bits)

3DES-ECB:

```go
CryptUtil.Des3ECBEncrypt(src, key, CryptUtil.PKCS7_PADDING)
CryptUtil.Des3ECBDecrypt(src, key, CryptUtil.PKCS7_PADDING)
```

3DES-CBC:

```go
CryptUtil.Des3CBCEncrypt(src, key, iv, CryptUtil.PKCS7_PADDING)
CryptUtil.Des3CBCDecrypt(src, key, iv, CryptUtil.PKCS7_PADDING)
```

### RSA

```go
CryptUtil.RSAGenerateKey(bits int, out io.Writer)
CryptUtil.RSAGeneratePublicKey(priKey []byte, out io.Writer)

CryptUtil.RSAEncrypt(src, pubKey []byte) ([]byte, error)
CryptUtil.RSADecrypt(src, priKey []byte) ([]byte, error)

CryptUtil.RSASign(src []byte, priKey []byte, hash crypto.Hash) ([]byte, error)
CryptUtil.RSAVerify(src, sign, pubKey []byte, hash crypto.Hash) error
```

### HMAC

```
// Sha1 计算字符串的Sha1哈希值
Sha1(str string) []byte

// Sha256 计算字符串的Sha1哈希值
Sha256(str string) []byte


// Hmac 使用HMAC方法计算字符串的sha1哈希值
HmacSha1(key string, data string) []byte

// HmacSha1ToString使用HMAC方法计算字符串的sha1哈希，输出小写十六进制
HmacSha1ToString(key string, data string) string

// HmacSha256使用HMAC方法计算字符串的sha256哈希值
HmacSha256(key string, data string) []byte

// HmacSha256ToString使用HMAC方法计算字符串的sha256哈希，输出小写十六进制
HmacSha256ToString(key string, data string) string
```

### BCRYPT

```go
// GenerateBcryptPassword 生成加密密码
GenerateBcryptPassword(password string) ([]byte, error)
// CompareHashAndPassword 验证加密密码和新输入验证
CompareHashAndPassword(hashPassword string, password string) (bool, error)
```

### CTR

```go
source := "hello world"
fmt.Println("原字符：", source)
key := "1234567812345678"
encryptCode, _ := AesCtrCrypt([]byte(source), []byte(key))
fmt.Println("密文：", string(encryptCode))
decryptCode, _ := AesCtrCrypt(encryptCode, []byte(key))
fmt.Println("解密：", string(decryptCode))

```