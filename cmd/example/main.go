package main

import (
	"fmt"
	lab1 "github.com/drunyaD/design-lab-1"
)

func main() {
	// TODO: Get input from the command line, handle errors.
	res, err := lab1.PrefixCalculate("+ 5 * - 4 2 3")
	fmt.Println(res,err)
}
