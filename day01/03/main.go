package main

import (
	"log"
	"unsafe"
)

func main() {
	var x int8
	var y int16
	var z int64
	var f int
	log.Println(unsafe.Sizeof(x))
	log.Println(unsafe.Sizeof(y))
	log.Println(unsafe.Sizeof(z))
	log.Println(unsafe.Sizeof(f))
}
