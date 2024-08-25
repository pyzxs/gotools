package ArrUtil

import "sort"

func SortArray(arr []interface{}, less func(i, j int) bool) {
	sort.Slice(arr, less)
}

func SortDesc(arr []interface{}, less func(i, j int) bool) {
	sort.Slice(arr, func(i, j int) bool {
		return less(j, i)
	})
}

func SortRecursive(arr []interface{}) {
	sort.SliceStable(arr, func(i, j int) bool {
		return Compare(arr[i], arr[j])
	})
}

// Compare 比较两个 interface{} 类型的值
func Compare(a, b interface{}) bool {
	switch a := a.(type) {
	case float64:
		if b, ok := b.(float64); ok {
			return a < b
		}
	case int:
		if b, ok := b.(int); ok {
			return a < b
		}
	case string:
		if b, ok := b.(string); ok {
			return a < b
		}
	case []interface{}:
		if b, ok := b.([]interface{}); ok {
			// 递归排序子切片
			SortRecursive(a)
			SortRecursive(b)
			return CompareSlice(a, b)
		}
	case map[string]interface{}:
		if b, ok := b.(map[string]interface{}); ok {
			// 对 map 的键值对进行排序
			sortedA := SortMap(a)
			sortedB := SortMap(b)
			return CompareSlice(sortedA, sortedB)
		}
	}
	return false
}

// CompareSlice 比较两个切片
func CompareSlice(a, b []interface{}) bool {
	length := len(a)
	if len(b) < length {
		length = len(b)
	}
	for i := 0; i < length; i++ {
		if !Compare(a[i], b[i]) {
			return false
		}
	}
	return len(a) < len(b)
}
