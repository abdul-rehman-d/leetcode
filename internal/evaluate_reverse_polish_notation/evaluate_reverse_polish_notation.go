package evaluate_reverse_polish_notation

import (
	"strconv"
)

func EvaluateReversePolishNotation(tokens []string) int {
	var recurse func() int
	recurse = func() int {
		popped := tokens[len(tokens)-1]
		tokens = tokens[:len(tokens)-1]

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
		return op(op2, op1)
	}

	return recurse()
}
