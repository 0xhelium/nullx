package nullx

import (
	"strings"
	"unicode/utf8"
)

func strlen(str string) int {
	return utf8.RuneCountInString(str)
}

func strrpt(str string, rpt int) string {
	return strings.Repeat(str, rpt)
}
