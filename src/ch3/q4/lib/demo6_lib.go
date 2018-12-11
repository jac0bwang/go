package lib

import (
	"os"
	in "ch3/q4/lib/internal"
)

func Hello(name string) {
	in.Hello(os.Stdout, name)
}
