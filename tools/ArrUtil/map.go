package ArrUtil

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"

	"github.com/pyzxs/gotools/convert"
)

// Collapse 将多个map合并成一个
func Collapse(arrs []map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for i, arr := range arrs {
		for k, v := range arr {
			result[convert.IntToStr(i)+"_"+k] = v
		}
	}
	return result
}


func Divide(m map[string]interface{}) ([]string, []interface{}) {
	keys := []string{}
	values := []interface{}{}
	for k, v := range m {
		keys = append(keys, k)
		values = append(values, v)
	}
	return keys, values
}

func Dot(m map[string]interface{}, prefix string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		switch v := v.(type) {
		case map[string]interface{}:
			nested := Dot(v, prefix+k+".")
			for nk, nv := range nested {
				result[nk] = nv
			}
		default:
			result[prefix+k] = v
		}
	}
	return result
}

func Except(m map[string]interface{}, keysToExclude []string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		exclude := false
		for _, key := range keysToExclude {
			if k == key {
				exclude = true
				break
			}
		}
		if !exclude {
			result[k] = v
		}
	}
	return result
}


func First(arr []interface{}) interface{} {
	if len(arr) > 0 {
		return arr[0]
	}
	return nil
}

func Flatten(arrs ...interface{}) []interface{} {
	result := []interface{}{}
	for _, item := range arrs {
		switch item := item.(type) {
		case []interface{}:
			result = append(result, Flatten(item...)...)
		default:
			result = append(result, item)
		}
	}
	return result
}


func Get(m map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if v, exists := m[key]; exists {
		return v
	}
	return defaultValue
}


// 判断是否存在
func HasAny(m map[string]interface{}, keys []string) bool {
	for _, key := range keys {
		if _, exists := m[key]; exists {
			return true
		}
	}
	return false
}



// 通过函数处理
func MapArray(arr []interface{}, fn func(interface{}) interface{}) []interface{} {
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

// 获取子集合
func Only(m map[string]interface{}, keys []string) map[string]interface{} {
	result := make(map[string]interface{})
	for _, key := range keys {
		if v, exists := m[key]; exists {
			result[key] = v
		}
	}
	return result
}

func Pluck(arr []map[string]interface{}, key string) []interface{} {
	result := []interface{}{}
	for _, item := range arr {
		if v, exists := item[key]; exists {
			result = append(result, v)
		}
	}
	return result
}

func Prepend(arr []interface{}, value interface{}) []interface{} {
	return append([]interface{}{value}, arr...)
}

func PrependKeysWith(m map[string]interface{}, prefix string) map[string]interface{} {
	result := make(map[string]interface{})
	for k, v := range m {
		result[prefix+k] = v
	}
	return result
}

func Pull(m map[string]interface{}, key string) interface{} {
	value := m[key]
	delete(m, key)
	return value
}

func Query(params map[string]interface{}) string {
	result := []string{}
	for k, v := range params {
		result = append(result, fmt.Sprintf("%s=%v", k, v))
	}
	return strings.Join(result, "&")
}

// 去除随机值
func Random(arr []interface{}) interface{} {
	if len(arr) == 0 {
		return nil
	}
	return arr[rand.Intn(len(arr))]
}

// 取出随机
func Shuffle(arr []interface{}) []interface{} {
	shuffled := make([]interface{}, len(arr))
	copy(shuffled, arr)
	rand.Shuffle(len(shuffled), func(i, j int) {
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})
	return shuffled
}

// SortMap 对 map 进行排序，返回排序后的键值对切片
func SortMap(m map[string]interface{}) []interface{} {
	var sorted []interface{}
	for k, v := range m {
		sorted = append(sorted, map[string]interface{}{k: v})
	}
	sort.SliceStable(sorted, func(i, j int) bool {
		return Compare(sorted[i], sorted[j])
	})
	return sorted
}
