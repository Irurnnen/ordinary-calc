package yetanothercalc

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

const disallowedSymbolsRegular = `[^0-9\.+\-*\/()^\s]`
const spacesRegular = `\s`

var operands = "+-*/^"
var priority = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "^": 3}

func Calc(expression string) (float64, error) {
	// Checking validity of expression
	if err := ValidateExpression(expression); err != nil {
		return 0, err
	}

	// Tokenize expression
	tokens := ParseExpression(expression)

	// Validate Tokens
	if err := ValidateTokens(tokens); err != nil {
		return 0, err
	}

	// Change to postfix
	postfixTokens := ToPostfix(tokens)

	// Calculate the expression
	result, err := EvalExpression(postfixTokens)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func ValidateExpression(expression string) error {
	// Check disallowed symbols
	re := regexp.MustCompile(disallowedSymbolsRegular)
	if re.MatchString(expression) {
		return ErrExtraCharacters
	}

	// Check correction of brackets
	var bracketBalance int
	for _, v := range expression {
		if v == '(' {
			bracketBalance++
		} else if v == ')' {
			bracketBalance--
			if bracketBalance < 0 {
				return ErrWrongBracketOrder
			}
		}
	}
	if bracketBalance != 0 {
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

	// Delete all space in expression
	expression = RemoveSpaces(expression)

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
	if len(number) != 0 {
		tokens = append(tokens, number)
	}
	return tokens
}

func ValidateTokens(tokens []string) error {
	// Check exists of expression
	if len(tokens) == 0 {
		return ErrEmptyExpression
	}
	// Check multiple operators or multiple numbers
	for i := 1; i < len(tokens); i++ {
		if IsOperand(tokens[i-1]) && IsOperand(tokens[i]) {
			return ErrMultipleOperands
		}
		if IsNumber(tokens[i-1]) && IsNumber(tokens[i]) {
			return ErrMultipleNumbers
		}
	}

	// Check operands at the beginning and end
	if IsOperand(tokens[0]) || IsOperand(tokens[len(tokens)-1]) {
		return ErrExtraOperands
	}

	return nil
}

func IsNumber(token string) bool {
	for _, v := range token {
		if v != '.' && (v < '0' || v > '9') {
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

func EvalExpression(tokens []string) (float64, error) {
	var stack []float64
	for _, token := range tokens {
		// If token is number
		if IsNumber(token) {
			num, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, ErrParseFloat
			}
			stack = append(stack, num)
			continue
		}
		// If token is operand
		a, b := stack[len(stack)-2], stack[len(stack)-1]
		stack = stack[:len(stack)-2]

		var result float64
		switch token {
		case "+":
			result = a + b
		case "-":
			result = a - b
		case "*":
			result = a * b
		case "/":
			if b == 0 {
				return 0, ErrZeroByDivision
			}
			result = a / b
		case "^":
			result = math.Pow(a, b)
		}

		stack = append(stack, result)
	}

	// Check size of stack
	if len(stack) == 0 {
		return 0, nil
	}

	return stack[0], nil
}
