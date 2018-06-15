package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	b, _ := ioutil.ReadAll(os.Stdin)
	fmt.Print(string(b))
}
