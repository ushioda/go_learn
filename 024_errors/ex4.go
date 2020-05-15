package main

import (
	"fmt"
	//"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Create("log.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println(err)
	}
	defer f.Close()
	log.SetOutput(f)

	f2, err := os.Open("no-file.txt")
	if err != nil {
		//fmt.Println("err happened", err)
		log.Println("err happened", err)
	}
	defer f2.Close()

	fmt.Println("check the log.txt file in the directory")

}
