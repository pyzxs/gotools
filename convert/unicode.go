package convert

import (
	"fmt"
	"strconv"
	"strings"
)

// StrToUnicode 将字符串转换为 Unicode 转义序列
func StrToUnicode(str string) string {
	var builder strings.Builder
	for _, r := range str {
		if r > 0x7F { // 只处理非 ASCII 字符
			builder.WriteString(fmt.Sprintf("\\u%04x", r))
		} else {
			builder.WriteRune(r)
		}
	}
	return builder.String()
}

// UnicodeToStr 将 Unicode 转义序列转换为原始字符串
func UnicodeToStr(unicodeStr string) (string, error) {
	var builder strings.Builder
	i := 0
	for i < len(unicodeStr) {
		if unicodeStr[i] == '\\' && i+5 < len(unicodeStr) && unicodeStr[i+1] == 'u' {
			code, err := strconv.ParseInt(unicodeStr[i+2:i+6], 16, 32)
			if err != nil {
				return "", err
			}
			builder.WriteRune(rune(code))
			i += 6
		} else {
			builder.WriteByte(unicodeStr[i])
			i++
		}
	}
	return builder.String(), nil
}
