package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	name := flag.String("l", "xxx.txt", "a file path")
	flag.Parse()
	f, err := os.Open(*name)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	body, _ := ioutil.ReadAll(f)
	fmt.Println(len(body))
}
