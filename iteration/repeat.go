package iteration

func Repeat(c rune) (rep string) {
	for range 5 {
		rep += string(c)
	}
	return
}

