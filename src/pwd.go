package main

import (
	"fmt"
	"os"
)

/* func main() {

	p, err := os.Getwd()
	if err != nil {	panic(err)	}
	fmt.Println(p)
}
*/
func pwd() {

	p, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}
