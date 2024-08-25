package convert

import (
	"fmt"
	"strings"
)

// NumberToWord 数字到英文转换的主函数
func NumberToWord(num float64) string {
	intPart := int(num)
	decimalPart := int((num - float64(intPart)) * 100) // 处理小数部分，假设精确到分

	// 获取英文表达
	intWord := convertIntToWord(intPart)
	decimalWord := convertIntToWord(decimalPart)

	// 构建最终的英文表达
	if decimalPart > 0 {
		return fmt.Sprintf("%s AND CENTS %s ONLY", intWord, decimalWord)
	}
	return fmt.Sprintf("%s ONLY", intWord)
}

// 将整数转换为英文
func convertIntToWord(num int) string {
	if num == 0 {
		return "ZERO"
	}

	// 单位数组
	thousands := []string{"", "THOUSAND", "MILLION", "BILLION"}

	var words []string

	if num >= 1000 {
		for i := 0; num > 0; i++ {
			if num%1000 != 0 {
				words = append(words, convertHundreds(num%1000), thousands[i])
			}
			num /= 1000
		}
	} else {
		words = append(words, convertHundreds(num))
	}

	return strings.Join(reverse(words), " ")
}

// 将千位以下的数转换为英文
func convertHundreds(num int) string {
	ones := []string{"", "ONE", "TWO", "THREE", "FOUR", "FIVE", "SIX", "SEVEN", "EIGHT", "NINE"}
	tens := []string{"", "", "TWENTY", "THIRTY", "FORTY", "FIFTY", "SIXTY", "SEVENTY", "EIGHTY", "NINETY"}

	var words []string

	if num >= 100 {
		words = append(words, ones[num/100], "HUNDRED")
		num %= 100
	}
	if num >= 20 {
		words = append(words, tens[num/10])
		num %= 10
	}
	if num > 0 {
		words = append(words, ones[num])
	} else if num == 0 && len(words) == 0 {
		words = append(words, "ZERO")
	}

	return strings.Join(words, " ")
}

// 反转字符串数组
func reverse(slice []string) []string {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	return slice
}
