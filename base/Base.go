package base

import (
	"bytes"
	"github.com/0xhelium/nullx/bit"
)

const BASE64_CHARSET = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"

func to6GroupBits(b []byte) bit.Bits {
	bits := make(bit.Bits, 0, len(b) * 6)
	for bi := 0; bi < len(b); bi++ {
		cb := b[bi]
		bits = append(bits, bit.ByteToBits(cb)[2:]...)
	}
	return bits
}

type Base64 []byte;
func (b64 Base64) Padding() int {
	var padding int = 0
	if b64[len(b64)-1] == byte('=') {
		padding += 1
	}
	if b64[len(b64)-2] == byte('=') {
		padding += 1
	}
	return padding
}
func (b64 Base64) Ascii() (string, error) {
	var b64i = make([]byte, len(b64))
	for i := 0; i < len(b64); i++ {
		b64i[i] = byte(bytes.IndexByte([]byte(BASE64_CHARSET), b64[i]))
	}
	var sb64, padding = StripPadding(b64i)
	var bits bit.Bits = to6GroupBits(sb64)
	bits = bits[:len(bits)-2*padding]
	return bits.Ascii()
}
func StripPadding(b64 Base64) (Base64, int) {
	var padding = b64.Padding()
	return b64[:len(b64)-padding], padding
}
