package convert

import (
	"strconv"
)

// ToDBC 将全角字符转换为半角字符
func ToDBC(input string) string {
	result := ""
	for _, r := range input {
		if r >= '\uFF01' && r <= '\uFF60' || r >= '\uFFE0' && r <= '\uFFE6' {
			// Convert full-width ASCII characters to half-width
			result += string(r - 0xFEE0)
		} else if r == '\u3000' {
			// Convert full-width space to half-width space
			result += " "
		} else {
			// Leave other characters unchanged
			result += string(r)
		}
	}
	return result
}

// ToSBC 将半角字符转换为全角字符
func ToSBC(input string) string {
	result := ""
	for _, r := range input {
		if r >= '\u0020' && r <= '\u007E' {
			// Convert half-width ASCII characters to full-width
			result += string(r + 0xFEE0)
		} else if r == ' ' {
			result += strconv.Itoa(0x3000)
		} else {
			// Leave other characters unchanged
			result += string(r)
		}
	}
	return result
}
