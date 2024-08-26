package HashUtil

import (
	"fmt"
	"os"
	"testing"
)

var (
	data []byte
)

func setup() {
	data = []byte("hello world")
}

func teardown() {
	fmt.Println("After all tests")
}

func TestAdditive(t *testing.T) {
	t.Log(Additive(data))
}

func TestRotating(t *testing.T) {
	t.Log(Rotating(data))
}

func TestOneByOne(t *testing.T) {
	t.Log(OneByOne(data))
}

func TestBernstein(t *testing.T) {
	t.Log(Bernstein(data))
}

func TestBkdrHash(t *testing.T) {
	t.Log(BkdrHash(data))
}

func TestUniversal(t *testing.T) {
	t.Log(Universal(data, 12, 34, 9, 2))
}

//func TestZobrist(t *testing.T) {
//	t.Log(Zobrist([]byte("abc"), [][]uint32{{1, 2, 3}}))
//}

func TestFnvHash(t *testing.T) {
	t.Log(FnvHash(data))
}

func TestIntHash(t *testing.T) {
	t.Log(IntHash(23))
}

func TestRsHash(t *testing.T) {
	t.Log(RsHash(data))
}
func TestJsHash(t *testing.T) {
	t.Log(JsHash(data))
}

func TestElfHash(t *testing.T) {
	t.Log(ElfHash(data))
}

func TestSdbmHash(t *testing.T) {
	t.Log(SdbmHash(data))
}

func TestDjbHash(t *testing.T) {
	t.Log(DjbHash(data))
}

func TestDekHash(t *testing.T) {
	t.Log(DekHash(data))
}

func TestPJWHash(t *testing.T) {
	t.Log(PJWHash("thisiskey"))
}

func TestApHash(t *testing.T) {
	t.Log(ApHash(data))
}
func TestTianlHash(t *testing.T) {
	t.Log(TianlHash(data))
}

//go:generate go test -v
func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
