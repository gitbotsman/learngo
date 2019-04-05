package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	body, _ := ioutil.ReadAll(os.Stdin)
	//fmt.Printf("%s : %d", string(body), len(body))
	fmt.Fprintf(os.Stdout, "%d\n", len(body))
}
