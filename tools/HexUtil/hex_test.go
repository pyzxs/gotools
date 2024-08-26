package HexUtil

import "testing"

func TestHex(t *testing.T) {
	hex := EncodeHex("test hex")
	decodeHex, _ := DecodeHex(hex)
	t.Log(hex, decodeHex)
}
