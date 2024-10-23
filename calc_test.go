package yetanothercalc_test

import (
	"testing"

	yetanothercalc "github.com/Irurnnen/yet-another-calc"
)

// TODO: write tests
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
			name:     "Correct advanced expression",
			input:    "15/(7-(1+1))*3-(2+(1+1))*15/(7-(200+1)^2)3-(2+(1+1))(15/(7-(1+1))*3-(2+(1+1))+15/(7-(1+1))*3-(2+(1+1)))",
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
