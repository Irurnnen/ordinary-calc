package yetanothercalc_test

import (
	"testing"

	yetanothercalc "github.com/Irurnnen/yet-another-calc"
)

func TestValidateExpression(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		excepted error
	}{
		{
			name:     "Correct expression",
			input:    "323+4*2/1-5^2",
			excepted: nil,
		},
		{
			name:     "Correct expression with spaces",
			input:    "12 * 34 / 45 + 4654739767725 ^ 21312",
			excepted: nil,
		},
		{
			name:     "Correct expression with spaces and brackets",
			input:    "59 * (213 + 231) / (10856 * (123 + 101) - 945) / 785",
			excepted: nil,
		},
		{
			name:     "Advanced expression",
			input:    "15/(7-(1+1))*3-(2+(1+1))*15/(7-(200+1)^2)3-(2+(1+1))(15/(7-(1+1))*3-(2+(1+1))+15/(7-(1+1))*3-(2+(1+1)))",
			excepted: nil,
		},
		{
			name:     "Only brackets",
			input:    "((()))",
			excepted: nil,
		},
		{
			name:     "Decimal numbers",
			input:    "2.5 + 3.7",
			excepted: nil,
		},
		{
			name:     "Extra characters",
			input:    "2 + a - 3",
			excepted: yetanothercalc.ErrExtraCharacters,
		},
		{
			name:     "Unpaired bracket (only opened bracket)",
			input:    "2 + (3 * 4",
			excepted: yetanothercalc.ErrUnpairedBracket,
		},
		{
			name:     "Unpaired brackets (only closed bracket)",
			input:    "2) + 3 * 4",
			excepted: yetanothercalc.ErrWrongBracketOrder,
		},
		{
			name:     "Unpaired brackets (wrong order)",
			input:    "2) + 3 * (4",
			excepted: yetanothercalc.ErrWrongBracketOrder,
		},
		{
			name:     "Empty expression",
			input:    "",
			excepted: nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := yetanothercalc.ValidateExpression(tc.input)

			if got != tc.excepted {
				t.Errorf("ValidateExpression(%q): got %q, excepted %q", tc.input, got, tc.excepted)
			}
		})
	}
}

// TODO: write tests for Calc
// TODO: write tests for RemoveSpaces
// TODO: write tests for ParseExpression
// TODO: write tests for ValidateTokens
// TODO: write tests for IsNumber
// TODO: write tests for IsOperand
// TODO: write tests for ToPostfix
// TODO: write tests for EvalExpression
