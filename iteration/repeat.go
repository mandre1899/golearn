package iteration

import "strings"

func Repeat(c rune, times int) string {
	var ret strings.Builder
	for range times {
		ret.WriteRune(c)
	}
	return ret.String()
}
