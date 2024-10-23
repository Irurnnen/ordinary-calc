package yetanothercalc

import (
	"regexp"
	"strings"
)

const disallowedSymbolsRegular = `[^0-9\.+\-*\/()^\s]`
const spacesRegular = `\s`

var operands = "+-*/^"
var operators = operands + "()"
var priority = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "^": 3}

func Calc(expression string) (float64, error) {
	// Checking validity of expression
	if err := ValidateExpression(expression); err != nil {
		return 0, err
	}
	// Delete all space in expression
	expression = RemoveSpaces(expression)

	// Tokenize expression
	tokens := ParseExpression(expression)

	// Validate Tokens
	if err := ValidateTokens(tokens); err != nil {
		return 0, err
	}

	// TODO: calculate the expression
	return 0, nil
}

func ValidateExpression(expression string) error {
	// Check disallowed symbols
	re := regexp.MustCompile(disallowedSymbolsRegular)
	if re.MatchString(expression) {
		return ErrExtraCharacters
	}

	// Check correction of brackets
	if strings.Count(expression, "(") != strings.Count(expression, ")") {
		return ErrUnpairedBracket
	}

	return nil
}

func RemoveSpaces(expression string) string {
	re := regexp.MustCompile(spacesRegular)
	return re.ReplaceAllString(expression, "")
}

func ParseExpression(expression string) []string {
	var tokens []string
	var number string

	for _, character := range expression {
		if IsNumber(string(character)) {
			number += string(character)
			continue
		}
		if number != "" {
			tokens = append(tokens, number)
			number = ""
		}
		tokens = append(tokens, string(character))
	}
	return tokens
}

func ValidateTokens(tokens []string) error {
	// Check multiple operators or multiple numbers
	for i := 1; i < len(tokens); i++ {
		if IsOperand(tokens[i-1]) == IsOperand(tokens[i]) {
			return ErrMultipleOperands
		}
		if IsNumber(tokens[i-1]) == IsNumber((tokens[i])) {
			return ErrMultipleNumbers
		}
	}
	return nil
}

func IsNumber(token string) bool {
	for _, v := range token {
		if v != '.' || v < '0' || v > '9' {
			return false
		}
	}
	return true
}

func IsOperand(token string) bool {
	return strings.Contains(operands, token) && len(token) == 1
}

func ToPostfix(tokens []string) []string {
	var stack []string
	var output []string

	for _, token := range tokens {
		if IsNumber(token) {
			output = append(output, token)
			continue
		}
		switch token {
		case "(":
			stack = append(stack, token)
		case ")":
			for len(stack) != 0 && stack[len(stack)-1] != "(" {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
		default:
			for len(stack) != 0 && stack[len(stack)-1] != "(" && priority[token] <= priority[stack[len(stack)-1]] {
				output = append(output, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, token)
		}
	}
	for len(stack) != 0 {
		output = append(output, stack[len(stack)-1])
		stack = stack[:len(stack)-1]
	}
	return output
}
