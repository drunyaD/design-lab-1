package main

import (
	"fmt"
	lab1 "github.com/drunyaD/design-lab-1"
	"os"
	"strings"
)

func main() {
	res, err := lab1.PrefixCalculate(strings.Join(os.Args[1:], " "))
	if err == nil {
		fmt.Println(res)
	} else {
		fmt.Println(err)
	}
}
