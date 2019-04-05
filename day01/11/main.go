package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	infos, err := ioutil.ReadDir("../../day01")
	if err != nil {
		log.Fatal(err)
	}
	for _, info := range infos {
		if info.IsDir() {
			fmt.Println(info.Name())
		}
	}
}
