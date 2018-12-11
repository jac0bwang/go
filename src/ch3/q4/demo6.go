package main

import (
	"flag"
	"ch3/q4/lib"
	//in "ch3/q4/lib/internal" // 此行无法通过编译。
	//"os"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Parse()
	lib.Hello(name)
	//in.Hello(os.Stdout, name)
}
