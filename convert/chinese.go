package convert

import (
	"fmt"
	"strings"
)

// 中文数字和单位
var digitMap = map[int]string{
	0: "零", 1: "壹", 2: "贰", 3: "叁", 4: "肆", 5: "伍", 6: "陆", 7: "柒", 8: "捌", 9: "玖",
}
var unitMap = []string{"", "拾", "佰", "仟"}
var bigUnitMap = []string{"", "万", "亿", "兆"}

// DigitToChinese 将数字转换为中文大写金额
func DigitToChinese(amount float64) string {
	// 将金额分解为整数部分和小数部分
	intPart := int(amount)
	decimalPart := int((amount - float64(intPart)) * 100)

	// 转换整数部分
	intPartChinese := convertIntegerToChinese(intPart)
	// 转换小数部分
	decimalPartChinese := convertDecimalToChinese(decimalPart)

	// 生成完整的中文金额
	if decimalPart > 0 {
		return fmt.Sprintf("%s元%s", intPartChinese, decimalPartChinese)
	}
	return fmt.Sprintf("%s元整", intPartChinese)
}

// convertIntegerToChinese 将整数部分转换为中文大写
func convertIntegerToChinese(number int) string {
	if number == 0 {
		return digitMap[0]
	}

	var result strings.Builder
	bigUnitPos := 0

	for number > 0 {
		part := number % 10000
		if part != 0 {
			partChinese := convertSmallNumberToChinese(part)
			if bigUnitPos > 0 {
				result.WriteString(bigUnitMap[bigUnitPos])
			}
			result.WriteString(partChinese)
		}
		number /= 10000
		bigUnitPos++
	}

	return reverseChineseString(result.String())
}

// convertSmallNumberToChinese 将四位数字转换为中文大写
func convertSmallNumberToChinese(number int) string {
	var result strings.Builder
	unitPos := 0

	for number > 0 {
		digit := number % 10
		if digit != 0 {
			result.WriteString(digitMap[digit])
			result.WriteString(unitMap[unitPos])
		} else {
			if unitPos > 0 {
				result.WriteString(digitMap[0])
			}
		}
		number /= 10
		unitPos++
	}

	return reverseChineseString(result.String())
}

// convertDecimalToChinese 将小数部分转换为中文大写
func convertDecimalToChinese(number int) string {
	if number == 0 {
		return ""
	}

	var result strings.Builder
	jiao := number / 10
	fen := number % 10

	if jiao > 0 {
		result.WriteString(digitMap[jiao])
		result.WriteString("角")
	}
	if fen > 0 {
		result.WriteString(digitMap[fen])
		result.WriteString("分")
	}
	return result.String()
}

// reverseChineseString 反转中文字符串
func reverseChineseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
