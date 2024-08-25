package StrUtil

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 测试字符串是否为空
func TestEmpty(t *testing.T) {
	t.Log(Empty(""), Empty(" "))
}

// 测试随机字符串
func TestRand(t *testing.T) {
	t.Log(Rand(4), Rand(5), Rand(300))
}

// 截取查询字符串后面内容
func TestAfter(t *testing.T) {
	a := assert.New(t)
	lists := []struct {
		s      string
		search string
		want   string
	}{
		{"this is dog", "is", " is dog"},
		{"this is a is of tea", " is", " a is of tea"},
		{"this is a is of tea", "a is", " of tea"},
	}

	for _, m := range lists {
		a.Equal(After(m.s, m.search), m.want)
	}
}

// 截取查询字符串前面
func TestBefore(t *testing.T) {
	a := assert.New(t)
	lists := []struct {
		s      string
		search string
		want   string
	}{
		{"this is dog", "is", "th"},
		{"this is a is of tea", " is", "this"},
		{"this is a is of tea", "a is", "this is "},
	}

	for _, m := range lists {
		a.Equal(Before(m.s, m.search), m.want)
	}
}

func TestAfterLast(t *testing.T) {
	a := assert.New(t)
	lists := []struct {
		s      string
		search string
		want   string
	}{
		{"this is dog", "is", " dog"},
		{"this is a is of tea", "is", "of tea"},
		{"this is a is of tea", "is", " of tea"},
	}

	for _, m := range lists {
		a.Equal(AfterLast(m.s, m.search), m.want)
	}
}

// 测试ASCII
func TestAscii(t *testing.T) {
	for _, v := range Ascii("!abcdABCD") {
		t.Logf("%d = %[1]c = %[1]b = %[1]x = %[1]U", v)
	}
}

// 首字母大写
func TestTitle(t *testing.T) {
	a := assert.New(t)
	a.Equal(Title("hello,world"), "Hello,World")
}

// 截取查询字符串前面
func TestBetween(t *testing.T) {
	a := assert.New(t)
	lists := []struct {
		s     string
		start string
		end   string
		want  string
	}{
		{"this is dog", "this", "dog", " is "},
		{"this is a is of tea", " is", "is", " a "},
		{"this is a is of tea", "is a", "tea", " of "},
	}

	for _, m := range lists {
		a.Equal(Between(m.s, m.start, m.end), m.want)
	}
}

func TestFormat(t *testing.T) {
	a := assert.New(t)

	mock := []struct {
		re string
		ps []string
		s  string
	}{
		{"{}", []string{"hello", "world"}, "{}你好{}"},
		{":[a-z_]+", []string{"8:30", "9:40"}, "time is :start - :end"},
	}

	for i, s := range mock {
		format, err := Format(s.re, s.ps, s.s)
		if i == 0 {
			a.Nil(err)
			a.Equal(format, "hello你好world")
		} else {
			a.Nil(err)
		}

	}
}

func TestCamel(t *testing.T) {
	t.Log(Camel("hello-world"))
}

func TestKebab(t *testing.T) {
	list := []struct {
		p    string
		want string
	}{
		{"getPage", "get-page"},
		{"GetDeptById", "get-dept-by-id"},
	}

	for _, s := range list {
		assert.Equal(t, Kebab(s.p), s.want)
	}

}
