package ordinarycalc

import (
	"fmt"
	"regexp"
	"strings"
)

const allowedSymbolsRegular = `[^0-9+\-*\/()^\s]`
const spacesRegular = `\s`

var operators = "+-*/()^"
var priority = map[string]int{"+": 1, "-": 1, "*": 2, "/": 2, "^": 3}

func Calc(expression string) (float64, error) {
	// Checking validity of expression
	if err := ValidateExpression(expression); err != nil {
		return 0, err
	}
	// Delete all space in expression
	expression = RemoveSpaces(expression)
	// tokens, err := parseExpression(expression)
	// if err != nil {
	// 	return 0, err
	// }
	tokens := ParseExpression(expression)
	fmt.Println(tokens)
	fmt.Println(ToPostfix(tokens))
	// TODO: calculate the expression
	return 0, nil
}

func ValidateExpression(expression string) error {
	re := regexp.MustCompile(allowedSymbolsRegular)
	if re.MatchString(expression) {
		return ErrExtraCharacters
	}
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
		if !strings.ContainsRune(operators, character) {
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

func ToPostfix(tokens []string) []string {
	var stack []string
	var output []string

	for _, token := range tokens {
		if !strings.Contains(operators, token) {
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
