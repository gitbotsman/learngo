package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("b.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	// fmt.Fprintf(f, "%d * %d = %2d", 3, 4, 3*4)
	for i := 1; i <= 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Fprintf(f, "%d * %d = %2d ", j, i, i*j)
		}
		f.WriteString("\n")
	}
}
