package main

import(
	"fmt"
	"os"
	"bufio"
	"log"
	

)

func main() {

	s := os.Args[1]
	f, err := os.Open(s)
	if err != nil{
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err :=scanner.Err(); err != nil {
		log.Fatal(err)
	}
}