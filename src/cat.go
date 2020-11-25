package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	switch a := len(os.Args);{
	case a < 2:
		os.Exit(0)
	case a == 2:
		readfile(os.Args[1])
	default:
		concatenate(os.Args[1], os.Args[2])
	}
}
func concatenate(file1, file2 string) {
	f, err := os.OpenFile(file2,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	f1, err := ioutil.ReadFile(file1)
	if err != nil {
		panic(err)
	}
	if _, err := f.WriteString(string(f1) + "\n"); err != nil {
		log.Println(err)
	}
}

func readfile(f string) {
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
