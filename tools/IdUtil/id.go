package IdUtil

import (
	"github.com/bwmarrin/snowflake"
	"log"
)

// SnowFlaskID 获取雪花算饭ID
func SnowFlaskID(n int64) string {
	node, err := snowflake.NewNode(n)
	if err != nil {
		log.Fatalf("Failed to create node: %v", err)
	}

	// 生成唯一 ID
	id := node.Generate()

	return id.String()
}
