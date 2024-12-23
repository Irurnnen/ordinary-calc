package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/Irurnnen/ordinary-calc/internal/forms"
	"github.com/Irurnnen/ordinary-calc/internal/models"
)

func TestCalcHandlerErrors(t *testing.T) {
	type args struct {
		expression     string
		exceptedCode   int
		exceptedResult float64
		exceptedError  string
	}
	testsFail := []struct {
		name string
		args args
	}{
		{
			name: "Expression with disallowed symbols",
			args: args{
				expression:    "2+2_2*2",
				exceptedCode:  422,
				exceptedError: "Expression has extra characters",
			},
		},
		{
			name: "Expression with unpaired open bracket",
			args: args{
				expression:    "2 + 2 * (2",
				exceptedCode:  422,
				exceptedError: "Expression has unpaired brackets",
			},
		},
		{
			name: "Expression with wrong bracket order",
			args: args{
				expression:    "2 + 2 ))(( * 2",
				exceptedCode:  422,
				exceptedError: "Expression has wrong bracket order",
			},
		},
		{
			name: "Expression with multiple operands",
			args: args{
				expression:    "2 + 2 ** 2",
				exceptedCode:  422,
				exceptedError: "Expression has multiple operands",
			},
		},
		{
			name: "Expression with zero by division (clearly)",
			args: args{
				expression:    "(2 + 2 * 2) / 0",
				exceptedCode:  422,
				exceptedError: "Expression has zero by division",
			},
		},
		{
			name: "Expression with zero by division (inconspicuous)",
			args: args{
				expression:    "(2 + 2 * 2) / (2 - 2)",
				exceptedCode:  422,
				exceptedError: "Expression has zero by division",
			},
		},
		{
			name: "Expression with extra operands at the beginning (clearly)",
			args: args{
				expression:    "+ 2 + 2 * 2",
				exceptedCode:  422,
				exceptedError: "Expression has operand at the beginning or at the end",
			},
		},
		{
			name: "Expression with extra operands at the beginning (inconspicuous)",
			args: args{
				expression:    "(+ 2) + 2 * 2",
				exceptedCode:  422,
				exceptedError: "Expression has operand at the beginning or at the end",
			},
		},
		{
			name: "Expression with extra operands at the end (clearly)",
			args: args{
				expression:    "2 + 2 * 2 * ",
				exceptedCode:  422,
				exceptedError: "Expression has operand at the beginning or at the end",
			},
		},
		{
			name: "Expression with extra operands at the end (inconspicuous)",
			args: args{
				expression:    "2 + 2 * (2 *)",
				exceptedCode:  422,
				exceptedError: "Expression has operand at the beginning or at the end",
			},
		},
		{
			name: "Expression with multiple operands",
			args: args{
				expression:    "",
				exceptedCode:  422,
				exceptedError: "Expression is empty",
			},
		},
	}
	for _, tt := range testsFail {
		t.Run(tt.name, func(t *testing.T) {
			// Data preparation
			expression := forms.Expression{
				Expression: tt.args.expression,
			}
			body, _ := json.Marshal(expression)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			// Create recorder
			recorder := httptest.NewRecorder()

			// Run handler
			CalcHandler(recorder, req)

			// Check http code
			if recorder.Code != tt.args.exceptedCode {
				t.Errorf("excepted status code %d, got %d", tt.args.exceptedCode, recorder.Code)
			}

			// Check body error
			var httpError forms.HTTPError
			err := json.NewDecoder(recorder.Body).Decode(&httpError)
			if err != nil {
				t.Errorf("error while decode json: %s", recorder.Body.String())
			}
			if httpError.Error != tt.args.exceptedError {
				t.Errorf("excepted error %s, got %s", tt.args.exceptedError, httpError.Error)
			}
		})
	}

	testsSuccess := []struct {
		name string
		args args
	}{
		{
			name: "Normal expression",
			args: args{
				expression:     "1+1+2+3",
				exceptedCode:   200,
				exceptedResult: 7,
			},
		},
		{
			name: "Expression with plus",
			args: args{
				expression:     "9+11",
				exceptedCode:   200,
				exceptedResult: 20,
			},
		},
		{
			name: "Expression with minus",
			args: args{
				expression:     "1+1+2-3",
				exceptedCode:   200,
				exceptedResult: 1,
			},
		},
		{
			name: "Expression with multiply",
			args: args{
				expression:     "2*3",
				exceptedCode:   200,
				exceptedResult: 6,
			},
		},
		{
			name: "Expression with division",
			args: args{
				expression:     "100/5",
				exceptedCode:   200,
				exceptedResult: 20,
			},
		},
		{
			name: "Expression with brackets",
			args: args{
				expression:     "(2 + 2) / 2",
				exceptedCode:   200,
				exceptedResult: 2,
			},
		},
		{
			name: "Expression with Priority",
			args: args{
				expression:     "2 * (3 - 24 / 4) + 2",
				exceptedCode:   200,
				exceptedResult: -4,
			},
		},
		{
			name: "Expression with spaces at the begging and end",
			args: args{
				expression:     " 123+1+233+3 ",
				exceptedCode:   200,
				exceptedResult: 360,
			},
		},
		{
			name: "Expression with spaces between characters",
			args: args{
				expression:     "1 +  1 +  2 +   3",
				exceptedCode:   200,
				exceptedResult: 7,
			},
		},
		{
			name: "Expression with multiple spaces",
			args: args{
				expression:     "\t\n\n\n1\t+\t    1\n\n+2\t\t+ \n\n\n3\t",
				exceptedCode:   200,
				exceptedResult: 7,
			},
		},
		{
			name: "Advanced expression",
			args: args{
				expression:     "(((45+15)*2-30)/3+(25*4-50))*2+(120/4-5*(3+7))+((30-15)*3+8/4)*5+(12*(5+3)-(10/2))-(100/(4+1))+15",
				exceptedCode:   200,
				exceptedResult: 461,
			},
		},
	}
	for _, tt := range testsSuccess {
		t.Run(tt.name, func(t *testing.T) {
			// Data preparation
			expression := forms.Expression{
				Expression: tt.args.expression,
			}
			body, _ := json.Marshal(expression)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/calculate", bytes.NewReader(body))
			req.Header.Set("Content-Type", "application/json")
			// Create recorder
			recorder := httptest.NewRecorder()

			// Run handler
			CalcHandler(recorder, req)

			// Check http code
			if recorder.Code != tt.args.exceptedCode {
				t.Errorf("excepted status code %d, got %d", tt.args.exceptedCode, recorder.Code)
			}

			// Check body error
			var result models.Result
			err := json.NewDecoder(recorder.Body).Decode(&result)
			if err != nil {
				t.Errorf("error while decode json: %s", recorder.Body.String())
			}
			if result.Result != tt.args.exceptedResult {
				t.Errorf("excepted error %f, got %f", tt.args.exceptedResult, result.Result)
			}
		})
	}
}
