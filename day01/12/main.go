package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("../../day01")
	if err != nil {
		log.Fatal(err)
	}
	infos, _ := f.Readdir(-1)

	for _, info := range infos {
		if info.IsDir() {
			fmt.Println(info.Name())
		}
	}
}
