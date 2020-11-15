package main

import(
	"fmt"
	"os"
	"bufio"
	"log"
	

)
//TODOS
// open and read file
// concatenate the file

func main() {

	if len(os.Args) < 2 {
		os.Exit(0)
	}else{
		read_file(os.Args[1])
	}
	
 
	
}

func read_file(f string){
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
	if err :=scanner.Err(); err != nil {
		log.Fatal(err)
	}

}