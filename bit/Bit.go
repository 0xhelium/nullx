package bit

import (
	"fmt"
	"math"
	"bytes"
	"strconv"
)

type Bits []byte;

func (bits Bits) String() string {
	return bits.ToString(" ")
}

func (bits Bits) ToString(delimiter string) string {
	str := ""
	for i := 0; i < len(bits); i++ {
		str += strconv.Itoa(int(bits[i]))
		if (i + 1) % 8 == 0 {
			str += delimiter
		}
	}
	return str
}

func (bits Bits) Ascii() (string, error) {
	bs, err := BitsToBytes(bits)
	if err != nil {
		return "", err
	}
	return string(bs), nil
}

func BitsToByte(bits Bits) byte {
	var b byte = 0;
	for i := 0; i < len(bits); i++ {
		b += byte(int(bits[len(bits)-i-1]) * int(math.Pow(2, float64(i))))
	}
	return b
}

func BitsToBytes(b Bits) ([]byte, error) {
	if len(b) % 8 != 0 {
		return nil, fmt.Errorf("size must be a multiple of 8 to be converted to ascii")
	}
	var bs []byte = make([]byte, len(b)/8, len(b)/8)
	buf := bytes.NewBuffer(b)
	for i := 0; i < len(b)/8; i++ {
		onebyte := buf.Next(8)
		bs[i] = BitsToByte(onebyte)
	}
	return bs, nil
}

func ByteToBits(b byte) Bits {
	bits := make(Bits, 8, 8)
	for i := 0; i < 8; i++ {
		place := int(math.Pow(2, float64(7 - i)))
		gi := int(b) / place
		bits[i] = byte(gi)
		if gi == 1 {
			b -= byte(place)
		}
	}
	return bits
}

func ToBits(b []byte) Bits {
	bits := make(Bits, 0, len(b) * 8)
	for bi := 0; bi < len(b); bi++ {
		cb := b[bi]
		bits = append(bits, ByteToBits(cb)...)
	}
	return bits
}

func Atob(binaryStringRaw string) Bits {
	var binaryString = []byte(binaryStringRaw)
	var binary = make(Bits, len(binaryString))
	for i := 0; i < len(binaryString); i++ {
		bin, err := strconv.Atoi(string(binaryString[i]))
		if err != nil {
			panic(err)
		}
		binary[i] = byte(bin)
	}
	return binary
}
