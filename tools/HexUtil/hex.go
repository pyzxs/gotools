package HexUtil

import "encoding/hex"

// EncodeHex 将字符串编码为十六进制表示
func EncodeHex(s string) string {
	// 转换字符串为字节切片
	data := []byte(s)
	// 将字节切片编码为十六进制字符串
	return hex.EncodeToString(data)
}

// DecodeHex 将十六进制字符串解码为原始字符串
func DecodeHex(hexStr string) (string, error) {
	// 将十六进制字符串解码为字节切片
	data, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", err
	}
	// 将字节切片转换为字符串
	return string(data), nil
}
