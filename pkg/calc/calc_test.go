package calc

import (
	"reflect"
	"testing"
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
			excepted: ErrExtraCharacters,
		},
		{
			name:     "Unpaired bracket (only opened bracket)",
			input:    "2 + (3 * 4",
			excepted: ErrUnpairedBracket,
		},
		{
			name:     "Unpaired brackets (only closed bracket)",
			input:    "2) + 3 * 4",
			excepted: ErrWrongBracketOrder,
		},
		{
			name:     "Unpaired brackets (wrong order)",
			input:    "2) + 3 * (4",
			excepted: ErrWrongBracketOrder,
		},
		{
			name:     "Empty expression",
			input:    "",
			excepted: nil,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := ValidateExpression(tc.input)

			if got != tc.excepted {
				t.Errorf("ValidateExpression(%q): got %q, excepted %q", tc.input, got, tc.excepted)
			}
		})
	}
}

// TODO: write tests for Calc
func TestCalc(t *testing.T) {
	casesSuccess := []struct {
		name           string
		input          string
		exceptedResult float64
	}{
		{
			name:           "Normal expression",
			input:          "1+1",
			exceptedResult: 2,
		},
		{
			name:           "Expression with plus",
			input:          "9+3",
			exceptedResult: 12,
		},
		{
			name:           "Expression with minus",
			input:          "9-3",
			exceptedResult: 6,
		},
		{
			name:           "Expression with multiply",
			input:          "9*3",
			exceptedResult: 27,
		},
		{
			name:           "Expression with division",
			input:          "9/3",
			exceptedResult: 3,
		},
		{
			name:           "Expression with priority",
			input:          "2 + 2 * 2",
			exceptedResult: 6,
		},
		{
			name:           "Expression with brackets",
			input:          "(2+2)*2",
			exceptedResult: 8,
		},
		{
			name:           "Spaces at beginning and end",
			input:          " 123 + 456 + 789 ",
			exceptedResult: 1368,
		},
		{
			name:           "Multiple spaces between characters",
			input:          "2  + 3    *      4",
			exceptedResult: 14,
		},
	}
	for _, tc := range casesSuccess {
		t.Run(tc.name, func(t *testing.T) {
			got, err := Calc(tc.input)
			if err != nil {
				t.Errorf("successful case %s return error %q", tc.name, err)
			}

			if got != tc.exceptedResult {
				t.Errorf("Calc(%q): got %f, excepted %f", tc.input, got, tc.exceptedResult)
			}
		})
	}
	casesFail := []struct {
		name        string
		expression  string
		expectedErr error
	}{
		{
			name:        "simple",
			expression:  "1+1*",
			expectedErr: ErrExtraOperands,
		},
	}
	for _, tc := range casesFail {
		t.Run(tc.name, func(t *testing.T) {
			_, err := Calc(tc.expression)
			if err == nil {
				t.Errorf("fail case %s does not return error %q", tc.name, tc.expectedErr)
			}
			if err != tc.expectedErr {
				t.Errorf("Calc(%q): got error %q, expected error %q", tc.expression, err, tc.expectedErr)
			}
		})
	}
}

func TestRemoveSpaces(t *testing.T) {
	cases := []struct {
		name     string
		input    string
		excepted string
	}{
		{
			name:     "Normal expression",
			input:    "2 + 3 * (4 - 5) / 6",
			excepted: "2+3*(4-5)/6",
		},
		{
			name:     "No spaces",
			input:    "2+3*(4-5)/5^6",
			excepted: "2+3*(4-5)/5^6",
		},
		{
			name:     "Only space",
			input:    " ",
			excepted: "",
		},
		{
			name:     "Empty expression",
			input:    "",
			excepted: "",
		},
		{
			name:     "Spaces at beginning and end",
			input:    " 123 + 456 + 789 ",
			excepted: "123+456+789",
		},
		{
			name:     "Multiple spaces between characters",
			input:    "2  + 3    *      4",
			excepted: "2+3*4",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := RemoveSpaces(tc.input)

			if got != tc.excepted {
				t.Errorf("RemoveSpaces(%q): got %q, excepted %q", tc.input, got, tc.excepted)
			}
		})
	}
}

func TestParseExpression(t *testing.T) {
	type args struct {
		expression string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Correct expression",
			args: args{"323+4*2/1-5^2"},
			want: []string{"323", "+", "4", "*", "2", "/", "1", "-", "5", "^", "2"},
		},
		{
			name: "Correct expression with spaces",
			args: args{"12 * 34 / 45 + 4654739767725 ^ 21312"},
			want: []string{"12", "*", "34", "/", "45", "+", "4654739767725", "^", "21312"},
		},
		{
			name: "Correct expression with spaces and brackets",
			args: args{"59 * (213 + 231) / (10856 * (123 + 101) - 945) / 785"},
			want: []string{"59", "*", "(", "213", "+", "231", ")", "/", "(", "10856", "*", "(", "123", "+", "101", ")", "-", "945", ")", "/", "785"},
		},
		{
			name: "Advanced expression",
			args: args{"15/(7-(1+1))*3-(2+(1+1))*15/(7-(200+1)^2)3-(2+(1+1))*(15/(7-(1+1))*3-(2+(1+1))+15/(7-(1+1))*3-(2+(1+1)))"},
			want: []string{"15", "/", "(", "7", "-", "(", "1", "+", "1", ")", ")", "*", "3", "-", "(", "2", "+", "(", "1", "+", "1", ")", ")", "*", "15", "/", "(", "7", "-", "(", "200", "+", "1", ")", "^", "2", ")", "3", "-", "(", "2", "+", "(", "1", "+", "1", ")", ")", "*", "(", "15", "/", "(", "7", "-", "(", "1", "+", "1", ")", ")", "*", "3", "-", "(", "2", "+", "(", "1", "+", "1", ")", ")", "+", "15", "/", "(", "7", "-", "(", "1", "+", "1", ")", ")", "*", "3", "-", "(", "2", "+", "(", "1", "+", "1", ")", ")", ")"},
		},
		{
			name: "Only brackets",
			args: args{"((()))"},
			want: []string{"(", "(", "(", ")", ")", ")"},
		},
		{
			name: "Decimal numbers",
			args: args{"2.5 + 3.7"},
			want: []string{"2.5", "+", "3.7"},
		},
		{
			name: "Empty expression",
			args: args{""},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseExpression(tt.args.expression); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseExpression() = %v, want %v", got, tt.want)
			}
		})
	}
}
