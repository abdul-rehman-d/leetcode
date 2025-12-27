package valid_parentheses

func ValidParentheses(s string) bool {
	stack := []byte{}

loop:
	for _, ch := range s {
		op := '.'
		switch ch {
		case '(', '{', '[':
			stack = append(stack, byte(ch))
			continue loop
		case ')':
			op = '('
		case '}':
			op = '{'
		case ']':
			op = '['
		default:
			panic("not parentheses")
		}
		if len(stack) == 0 {
			return false
		}
		popped := stack[len(stack)-1]
		if popped != byte(op) {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

