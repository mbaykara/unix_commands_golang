package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)


func main() {
	a := os.Args
	b := os.Args[1:]
 parseArgs(a)
 parseArgs(b)
}

func parseArgs(arguments []string) {
	files, err := filepath.Glob("[^.]*")
	if err != nil{ log.Fatal(err) }
	
	if len(arguments) == 0 {
		for _, file := range files {
			fmt.Println(file)
		}
	}
	if len(arguments) == 1 {
		switch arg := arguments[0]; arg {
		case "-l":
			for _, file:= range files{
				info,_ := os.Stat(file)
				fmt.Println(info.Mode(), file )
			}
		case "-la":
			files,_ := filepath.Glob("*")
			for _, file:= range files{
				info,_ := os.Stat(file)
				fmt.Println(info.Mode(), file )
			}
		case "-R", "-lR", "-r", "-lr":
		   err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
 		   if err != nil {  return err   }
 		   fmt.Println(path)
 		   return nil
			})
			if err != nil { log.Println(err)}
		case "-h", "--help":
				fmt.Println("Help will be soon")
		case "-a":
			files, err := ioutil.ReadDir(".")
			if err != nil{	log.Fatal(err) }
			for _,file := range files {
				fmt.Println(file.Name())
			}
			fmt.Print("test")
		}
		
	}


}
