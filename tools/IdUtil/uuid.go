package IdUtil

import (
	"github.com/google/uuid"
)

// RandomUUID 生成的UUID是带-的字符串，类似于：a5c8a5e8-df2b-4706-bea4-08d0939410e3
func RandomUUID() string {
	return uuid.New().String()
}

// SimpleUUID 生成的是不带-的字符串，类似于：b17f24ff026d40949c85a24f4f375d42
func SimpleUUID() string {
	return RemoveDashes(uuid.New().String())
}

func RemoveDashes(uuidStr string) string {
	return string([]rune(uuidStr)[0:8]) +
		string([]rune(uuidStr)[9:13]) +
		string([]rune(uuidStr)[14:18]) +
		string([]rune(uuidStr)[19:23]) +
		string([]rune(uuidStr)[24:])
}
