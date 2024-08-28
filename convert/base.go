package convert

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// StrToInt64 字符串转为为int43
func StrToInt64(e string, defaultValue int64) (int64, error) {
	v, err := strconv.ParseInt(e, 10, 64)
	if err != nil {
		return defaultValue, err
	}

	return v, nil
}

// Int64ToStr 整型64位转化为字符串
func Int64ToStr(e int64) string {
	return strconv.FormatInt(e, 10)
}

// Float64ToStr 浮点转化为字符串
func Float64ToStr(e float64) string {
	return strconv.FormatFloat(e, 'E', -1, 64)
}

// StrToFloat64 字符串转换64位浮点数
func StrToFloat64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

// Float32ToStr 32位浮点数转为字符串
func Float32ToStr(e float32) string {
	return fmt.Sprintf("%f", e)
}

// StrToInt 字符串转化为int
func StrToInt(e string, defaultValue int) (int, error) {
	v, err := strconv.Atoi(e)
	if err != nil {
		return defaultValue, err
	}
	return v, nil
}

// IntToStr int转化为字符串
func IntToStr(s int) string {
	return strconv.Itoa(s)
}

// IdsStrToSlice 传递的"1,2,4"字符串转化为切片
func IdsStrToSlice(id string) []int {
	ids := make([]int, 0)
	for v := range strings.Split(id, ",") {
		ids = append(ids, v)
	}

	return ids
}

// 字符串转Boolean型
func StrToBoolean(str string) bool {
	value, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return value
}

// JsonToMap json转化为map
func JsonToMap(e string) (map[string]interface{}, error) {
	var dict map[string]interface{}
	if err := json.Unmarshal([]byte(e), &dict); err == nil {
		return dict, err
	} else {
		return nil, err
	}
}

// StructToJson 结构体转化为json
func StructToJson(e interface{}) (string, error) {
	if b, err := json.Marshal(e); err == nil {
		return string(b), err
	} else {
		return "", err
	}
}

// StructToMap 结构体转化为map
func StructToMap(data interface{}) (map[string]interface{}, error) {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	mapData := make(map[string]interface{})
	err = json.Unmarshal(dataBytes, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

// 强制转化
func ForceCovert(src any, dst interface{}) error {
	b, err := json.Marshal(src)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, dst)
	if err != nil {
		return err
	}
	return nil
}

// ToDate 将日期字符串转换为 time.Time 对象
func StrToDate(dateStr string) (time.Time, error) {
	// 定义日期格式
	layout := "2006-01-02"
	// 解析日期字符串
	return time.Parse(layout, dateStr)
}
