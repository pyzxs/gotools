package IdUtil

import (
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
)

// RandomUlid 唯一词典分类标识符
func RandomUlid() string {
	entropy := rand.New(rand.NewSource(time.Now().UnixNano()))
	ms := ulid.Timestamp(time.Now())
	u, _ := ulid.New(ms, entropy)
	return u.String()
}
