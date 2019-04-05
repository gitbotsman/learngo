package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {
	a := "你好世界"
	a1 := []rune(a)
	a1[0] = '不'
	a = string(a1)
	fmt.Println(a) //不好世界
	fmt.Println(3.1 / 2)
	b := "12.2"
	c, err := strconv.Atoi(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)
}
