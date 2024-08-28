package IdUtil

import (
	"strings"
	"testing"
	"time"

	"github.com/pyzxs/gotools/convert"
	"github.com/pyzxs/gotools/tools/ArrUtil"
)

func TestUnionId(t *testing.T) {
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()
	i := 0
	for c := range ticker.C {
		i++
		if i >= 10 {
			break
		}
		t.Log(c, SnowFlaskID(1), RandomUUID(), SimpleUUID(), RandomUlid())
	}
}

func TestHashIds(t *testing.T) {
	input := []int{1, 2, 3}
	encode := HashIdsEncode(input)
	decode := HashIdsDecode(encode)
	if ArrUtil.Compare(decode, input) {
		t.Errorf("Before: %q, after: %q", input, decode)
	}

	t.Log(encode, decode)
}

func FuzzHashIds(f *testing.F) {
	for _, ints := range []string{"1,2,3", "3,5,6"} {
		f.Add(ints)
	}
	f.Fuzz(func(t *testing.T, input string) {
		split := strings.Split(input, ",")
		inters := make([]int, len(split))
		for i, s := range split {
			inters[i], _ = convert.StrToInt(s, 0)
		}
		t.Log(inters)
		encode := HashIdsEncode(inters)
		decode := HashIdsDecode(encode)
		if !ArrUtil.Compare(decode, input) {
			t.Errorf("Before: %q, after: %q", input, decode)
		}
		t.Log(encode, decode)
	})
}
