package lab1

import (
	"fmt"
	"strings"
	"strconv"
)

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func isNumeric(s string) bool {
    _, err := strconv.ParseInt(s, 10, 64)
    return err == nil
}

// TODO: document this function.
// Prefix form calculates
func PrefixCalculate(input string) (string	, error) {
	if input == "" {
		return "", fmt.Errorf("input should not be empty")
	}

	//res := 0
	splitedString := strings.Split(input, " ")
	signs := []string{"+", "-", "*", "/", "^"}

	for _, symbol := range splitedString {
		if (contains(signs, symbol) == false && isNumeric(symbol) == false) {
			return "", fmt.Errorf("failed on parsing : invalid input")
		}
		
	}
	return "TODO", fmt.Errorf("TODO")
}
