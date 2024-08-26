package EscapeUtil

import "testing"

func TestEscape(t *testing.T) {
	e := Escape("你好,abc")
	t.Log(e)
	t.Log(Unescape(e))
	t.Log(SafeUnescape(e))
}
