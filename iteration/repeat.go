package iteration

import "strings"

const repeatCount = 5

func Repeat(c rune) string {
	var ret strings.Builder
	for range repeatCount {
		ret.WriteRune(c)
	}
	return ret.String()
}
