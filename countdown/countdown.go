package countdown

import (
	"fmt"
	"io"
)

const counter = 3
const finalWord = "Go!"

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := range counter {
		if i < counter {
			fmt.Fprintf(writer, "%d\n", counter - i)
		}
		sleeper.Sleep()
	}
	sleeper.Sleep()
	fmt.Fprintf(writer, finalWord + "\n")
}
