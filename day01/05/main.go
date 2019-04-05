package main

import "fmt"

func main() {
	var a = 1
	var b = 1
	var s = 1
	for b <= 100 {
		a, b = b, a+b //到最后b 是大于100才退出，a是小于100
		s += a
	}
	fmt.Println(s)
}
