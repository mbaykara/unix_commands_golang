package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	efficient()
	lessefficient()
}

func efficient() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func lessefficient() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
