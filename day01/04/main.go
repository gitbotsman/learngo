package main

import "fmt"

func main() {
	a := "abc"
	// aByte := []byte(a)
	// aByte[0] = '1'
	// a = string(aByte)
	a = "1" + a[1:]

	fmt.Println(a)
}
