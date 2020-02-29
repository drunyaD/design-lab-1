package lab1

import (
	"fmt"
	"strconv"
	"strings"
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
func PrefixCalculate(input string) (string, error) {
	if input == "" {
		return "", fmt.Errorf("input should not be empty")
	}

	splitedString := strings.Split(input, " ")
	signs := []string{"+", "-", "*", "/", "^"}
	var stack []string
	for i := 0; i < len(splitedString); i++ {
		symbol := splitedString[i]
		if contains(signs, symbol) == false && isNumeric(symbol) == false {
			return "", fmt.Errorf("failed on parsing : invalid input")
		}
		stack = append(stack, symbol)

		for len(stack) >= 3 && isNumeric(stack[len(stack)-1]) && isNumeric(stack[len(stack)-2]) && contains(signs, stack[len(stack)-3]) {

			switch sign := stack[len(stack)-3]; sign {
			case "+":
				arg1, _ := strconv.ParseInt(stack[len(stack)-2], 10, 64)
				arg2, _ := strconv.ParseInt(stack[len(stack)-1], 10, 64)
				res := arg1 + arg2
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, strconv.FormatInt(res, 10))
			case "-":
				arg1, _ := strconv.ParseInt(stack[len(stack)-2], 10, 64)
				arg2, _ := strconv.ParseInt(stack[len(stack)-1], 10, 64)
				res := arg1 - arg2
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, strconv.FormatInt(res, 10))
			case "*":
				arg1, _ := strconv.ParseInt(stack[len(stack)-2], 10, 64)
				arg2, _ := strconv.ParseInt(stack[len(stack)-1], 10, 64)
				res := arg1 * arg2
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, strconv.FormatInt(res, 10))
			case "/":
				arg1, _ := strconv.ParseInt(stack[len(stack)-2], 10, 64)
				arg2, _ := strconv.ParseInt(stack[len(stack)-1], 10, 64)
				res := arg1 / arg2
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, strconv.FormatInt(res, 10))
			case "^":
				arg1, _ := strconv.ParseInt(stack[len(stack)-2], 10, 64)
				arg2, _ := strconv.ParseInt(stack[len(stack)-1], 10, 64)
				res := arg1 / arg2
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = stack[:len(stack)-1]
				stack = append(stack, strconv.FormatInt(res, 10))
			}
		}
	}

	if (len(stack) == 1) && isNumeric(stack[len(stack)-1]) {
		return stack[len(stack)-1], nil
	} else {
		return "", fmt.Errorf("failed on calculating : invalid input")
	}
}
