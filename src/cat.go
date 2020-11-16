package main

import(
	"fmt"
	"os"
	"bufio"
	"log"
	"io/ioutil"

)

func main() {

 	if len(os.Args) < 2 {
		os.Exit(0)
	}else if len(os.Args) == 2{
		read_file(os.Args[1])
	} else {
		concatenate(os.Args[1], os.Args[2])
	}
	
}
func concatenate(file1, file2 string){

	f, err := os.OpenFile(file2,
	os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	//read file1
	f1, err := ioutil.ReadFile(file1)
    if err != nil {
		panic(err)
	}

	if _, err := f.WriteString( string(f1)+"\n"); err != nil {
		log.Println(err)
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