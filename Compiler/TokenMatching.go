package main

func checkValidToken(token string) bool {
	brackets := map[string]bool {
	    "[": true,
	    "]": true,
	    "{": true,
	    "}": true,
	    "(": true,
    	")": true,
	}

	operators := map[string]bool {
		"+": true,
		"-": true,
		"/": true,
		"=": true,
	}

	_, isBracket := brackets[token]
	_, isOperator := operators[token]

	if isBracket || isOperator {
		return true
	}

	return false
}

func checkSpecialCharacter(token string) bool {
	whitespace := map[string]bool {
		" ": true,
		"\n": true,
		"\r": true,
	}

	_, exists := whitespace[token]
	return exists
}