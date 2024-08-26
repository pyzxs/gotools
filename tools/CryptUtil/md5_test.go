package CryptUtil

import (
	"encoding/base64"
	"encoding/hex"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	src := "apple"
	dst := Md5(src)
	assert.Equal(t, hex.EncodeToString(dst), "1f3870be274f6c49b3e31a0c6728957f")

	contentMD5 := base64.StdEncoding.EncodeToString([]byte(dst))
	t.Log(contentMD5)
}
