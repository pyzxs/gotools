package StrUtil

import (
	"fmt"
	"math/rand"
	"regexp"
	"strings"
	"unicode"
)

// Empty 判断一个字符串是否为空
func Empty(s string) bool {
	return len(s) == 0
}

// Rand 生成随机字符串
func Rand(length int) string {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	if length > len(letters) {
		length = len(letters)
	}
	result := make([]byte, length)
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

// After 返回字符串中指定值之后的所有内容
func After(s, search string) string {
	index := strings.Index(s, search)
	if index == -1 {
		return s
	}
	return s[index+len(search):]
}

// AfterLast 方法返回字符串中指定值最后一次出现后的所有内容
func AfterLast(s, search string) string {
	index := strings.LastIndex(s, search)
	if index == -1 {
		return s
	}
	return s[index+len(search):]
}

// Ascii 字符串转换为 ASCII
func Ascii(s string) []int {
	var asciiValues []int
	for _, char := range s {
		asciiValues = append(asciiValues, int(char))
	}
	return asciiValues
}

// Before 法返回字符串中指定值之前的所有内容
func Before(s, search string) string {
	index := strings.Index(s, search)
	if index == -1 {
		return s
	}
	return s[:index]
}

// BeforeLast 方法返回字符串中指定值最后一次出现前的所有内容
func BeforeLast(s, search string) string {
	index := strings.LastIndex(s, search)
	if index == -1 {
		return s
	}
	return s[:index]
}

// Between 法返回字符串在指定两个值之间的内容
func Between(s, start, end string) string {
	startIndex := strings.Index(s, start)
	if startIndex == -1 {
		return ""
	}
	startIndex += len(start)
	endIndex := strings.Index(s[startIndex:], end)
	if endIndex == -1 {
		return ""
	}
	return s[startIndex : startIndex+endIndex]
}

// Format 替换字符串中的正则匹配
func Format(pattern string, replacements []string, subject string) (string, error) {
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", fmt.Errorf("invalid pattern: %v", err)
	}

	result := re.ReplaceAllStringFunc(subject, func(match string) string {
		if len(replacements) > 0 {
			replacement := replacements[0]
			replacements = replacements[1:]
			return replacement
		}
		return match
	})

	return result, nil
}

// Camel 法将给定的字符串转换为 camelCase
func Camel(s string) string {
	words := strings.FieldsFunc(s, unicode.IsPunct)
	for i := range words {
		if i > 0 {
			words[i] = strings.Title(words[i])
		}
	}
	return strings.Join(words, "")
}

// Contains 字符串中是否含有子串
func Contains(s, substr string) bool {
	return strings.Contains(s, substr)
}

// EndsWith 判断指定字符串是否以另一指定字符串结尾
func EndsWith(s, substr string) bool {
	return strings.HasSuffix(s, substr)
}

// Lower 字符串转换为小写
func Lower(s string) string {
	return strings.ToLower(s)
}

// Upper 字符串转化为大写
func Upper(s string) string {
	return strings.ToUpper(s)
}

// Kebab 驼峰的函数名或者字符串转换成-
func Kebab(s string) string {
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '-')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// Snake 驼峰的函数名或者字符串转换成_
func Snake(s string) string {
	s = Lower(s)
	var result []rune
	for i, r := range s {
		if unicode.IsUpper(r) {
			if i > 0 {
				result = append(result, '_')
			}
			result = append(result, unicode.ToLower(r))
		} else {
			result = append(result, r)
		}
	}
	return string(result)
}

// Title 给定的字符串转换为Title Case
func Title(s string) string {
	return strings.Title(s)
}
