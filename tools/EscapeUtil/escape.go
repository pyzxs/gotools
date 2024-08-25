package EscapeUtil

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Escape Escape编码（Unicode），ASCII字母、数字和特定标点符号不进行编码
func Escape(input string) string {
	var builder strings.Builder
	for _, r := range input {
		if ('a' <= r && r <= 'z') || ('A' <= r && r <= 'Z') || ('0' <= r && r <= '9') ||
			r == '*' || r == '@' || r == '-' || r == '_' || r == '+' || r == '.' || r == '/' {
			builder.WriteRune(r)
		} else {
			builder.WriteString(fmt.Sprintf("\\u%04X", r))
		}
	}
	return builder.String()
}

// Unescape 解码
func Unescape(input string) (string, error) {
	re := regexp.MustCompile(`\\u([0-9A-Fa-f]{4})`)
	result := re.ReplaceAllStringFunc(input, func(s string) string {
		hex := s[2:] // 去掉前缀\\u
		code, _ := strconv.ParseInt(hex, 16, 32)
		return string(rune(code))
	})
	return result, nil
}

// SafeUnescape 安全解码，如果文本没有被escape则返回原文
func SafeUnescape(input string) string {
	// 尝试进行解码
	decoded, err := Unescape(input)
	if err != nil || !strings.Contains(input, "\\u") {
		return input
	}
	return decoded
}
