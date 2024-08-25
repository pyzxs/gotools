package ArrUtil

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
)

// IsNULL 判断是否为NULL
func IsNULL(v map[interface{}]interface{}) bool {
	return v == nil
}

// IsEmpty 判断是否为空
func IsEmpty(v map[interface{}]interface{}) bool {
	return len(v) == 0
}

// Add 向map中添加一个值
func Add(m map[string]interface{}, key string, value interface{}) {
	m[key] = value
}

// Collapse 将多个map合并成一个
func Collapse(arrs []map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{})
	for _, arr := range arrs {
		for k, v := range arr {
			result[k] = v
		}
	}
	return result
}

func CrossJoin(arrs ...[]interface{}) [][]interface{} {
	if len(arrs) == 0 {
		return nil
	}
	result := [][]interface{}{}
	for _, v := range arrs[0] {
		if len(arrs) == 1 {
			result = append(result, []interface{}{v})
		} else {
			rest := CrossJoin(arrs[1:]...)
			for _, r := range rest {
				result = append(result, append([]interface{}{v}, r...))
			}
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

func Exists(m map[string]interface{}, key string) bool {
	_, exists := m[key]
	return exists
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

func Forget(m map[string]interface{}, key string) {
	delete(m, key)
}

func Get(m map[string]interface{}, key string, defaultValue interface{}) interface{} {
	if v, exists := m[key]; exists {
		return v
	}
	return defaultValue
}

func Has(m map[string]interface{}, key string) bool {
	_, exists := m[key]
	return exists
}

func HasAny(m map[string]interface{}, keys []string) bool {
	for _, key := range keys {
		if _, exists := m[key]; exists {
			return true
		}
	}
	return false
}

func IsAssoc(m map[string]interface{}) bool {
	if len(m) == 0 {
		return false
	}
	for _, k := range m {
		if _, ok := k.(int); ok {
			return false
		}
	}
	return true
}

func last(arr []interface{}) interface{} {
	if len(arr) > 0 {
		return arr[len(arr)-1]
	}
	return nil
}

func MapArray(arr []interface{}, fn func(interface{}) interface{}) []interface{} {
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		result[i] = fn(v)
	}
	return result
}

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

func Random(arr []interface{}) interface{} {
	if len(arr) == 0 {
		return nil
	}
	return arr[rand.Intn(len(arr))]
}

func Set(m map[string]interface{}, key string, value interface{}) {
	m[key] = value
}

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
