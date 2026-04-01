var openBracketsByClosing = map[rune]rune{
	')': '(',
	']': '[',
	'}': '{',
}

func isValid(s string) bool {
	if len(s) < 2 {
		return false
	}

	stack := make([]rune, 0, len(s))

	// Enqueue open brackets into the stack iff they are an open bracket
	for _, curr := range s {
		if !isClosingBracket(curr) {
			stack = append(stack, curr)
			continue 
		}
		// Case where we have closing brackets with no matching open brackets
		if len(stack) == 0 {
			return false
		}

		// If the bracket is not an open bracket, is an open bracket
		// Pop from the stack, expecting the corresponding open bracket to the current closing bracket
		pop := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		expected, ok := fetchCorrespondingOpenBracket(curr)
		if !ok {
			return false
		}
		if pop != expected {
			return false
		}
	}

	return len(stack) == 0
}

func isClosingBracket(r rune) bool {
	_, ok := openBracketsByClosing[r]
	return ok
}

func fetchCorrespondingOpenBracket(r rune) (rune, bool) {
	v, ok := openBracketsByClosing[r]
	return v, ok 
}
