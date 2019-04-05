package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("a.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.WriteString("12121212")
}
