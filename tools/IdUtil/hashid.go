package IdUtil

import "github.com/speps/go-hashids/v2"

const (
	MIN_LENGTH   = 5
	DEFAUlT_SALE = "gotools"
)

var (
	idData = &hashids.HashIDData{
		Alphabet:  hashids.DefaultAlphabet,
		MinLength: MIN_LENGTH,
		Salt:      DEFAUlT_SALE,
	}
)

// HashIdsEncode 切片生成随机短串
func HashIdsEncode(number []int) string {
	id, _ := hashids.NewWithData(idData)
	encode, _ := id.Encode(number)
	return encode
}

// HashIdsDecode 解析字符串获取还原数据
func HashIdsDecode(input string) []int {
	en, _ := hashids.NewWithData(idData)
	withError, _ := en.DecodeWithError(input)
	return withError
}
