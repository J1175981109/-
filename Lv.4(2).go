package main

import (
	"fmt"
	"log"
	"os"
)

func main(){
	file,err := os.Open("proverb.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	bytes := make([]byte,65)
	br,err := file.Read(bytes)
	if err != nil{
		log.Fatal(err)
	}
	log.Printf("Number of Bytes read: %d\n",br)
	fmt.Println(string(bytes))
}
