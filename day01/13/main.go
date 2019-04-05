package main

import (
	"fmt"
)

func print1() *int {
	var x int
	return &x
}

func main() {
	p := print1()
	fmt.Println(*p)
}
