package main

import "fmt"

func main() {
	str := "Go爱好者"
	for i, c := range str {
		fmt.Printf("%d: %q [% x] len = %d\n", i, c, []byte(string(c)), len(string(c)))
	}
}
