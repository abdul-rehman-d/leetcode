package evaluate_reverse_polish_notation

import (
	"fmt"
	"strconv"
)

func EvaluateReversePolishNotation(tokens []string) int {
	out := 0

	fmt.Println()
	fmt.Println()

	var recurse func() int
	recurse = func() int {
		popped := tokens[len(tokens)-1]
		tokens = tokens[:len(tokens)-1]

		fmt.Println(popped, tokens)

		var op func(x, y int) int

		switch popped {
		case "+":
			op = func(x, y int) int {
				return x + y
			}
		case "-":
			op = func(x, y int) int {
				return x - y
			}
		case "/":
			op = func(x, y int) int {
				return x / y
			}
		case "*":
			op = func(x, y int) int {
				return x * y
			}
		default:
			i, _ := strconv.Atoi(popped)
			return i
		}
		op1 := recurse()
		op2 := recurse()
		fmt.Println(op1, op2, popped)
		return op(op2, op1)
	}

	out = recurse()

	return out
}
