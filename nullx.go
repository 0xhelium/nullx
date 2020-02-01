package nullx

import (
	"strings"
	"unicode/utf8"
)

func Strlen(str string) int {
	return utf8.RuneCountInString(str)
}

func Strrpt(str string, rpt int) string {
	return strings.Repeat(str, rpt)
}
