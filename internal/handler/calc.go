package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Irurnnen/ordinary-calc/internal/forms"
	"github.com/Irurnnen/ordinary-calc/internal/models"
	"github.com/Irurnnen/ordinary-calc/pkg/calc"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	// Get data from request
	var expression models.Expression

	err := json.NewDecoder(r.Body).Decode(&expression)
	if err != nil {
		ErrorJSONHandler(w, http.StatusBadRequest, forms.HTTPError{Error: "Provided data is invalid"})
		return
	}

	// Calculate the expression
	result, err := calc.Calc(expression.Expression)

	// Process errors
	switch err {
	case nil:
		break
	case calc.ErrExtraCharacters:
		ErrorJSONHandler(w, http.StatusUnprocessableEntity, forms.HTTPError{Error: "Expression has extra characters"})
		return
	case calc.ErrUnpairedBracket:
		ErrorJSONHandler(w, http.StatusBadRequest, forms.HTTPError{Error: "Expression has unpaired brackets"})
		return
	case calc.ErrWrongBracketOrder:
		ErrorJSONHandler(w, http.StatusBadRequest, forms.HTTPError{Error: "Expression has wrong bracket order"})
		return
	case calc.ErrMultipleOperands:
		ErrorJSONHandler(w, http.StatusBadRequest, forms.HTTPError{Error: "Expression has multiple operands"})
		return
	case calc.ErrMultipleNumbers:
		ErrorJSONHandler(w, http.StatusBadRequest, forms.HTTPError{Error: "Expression has multiple sequential numbers"})
		return
	case calc.ErrZeroByDivision:
		ErrorJSONHandler(w, http.StatusBadRequest, forms.HTTPError{Error: "Expression has zero by division"})
		return
	case calc.ErrExtraOperands:
		ErrorJSONHandler(w, http.StatusBadRequest, forms.HTTPError{Error: "Expression has at the beginning or at the end"})
		return
	case calc.ErrEmptyExpression:
		ErrorJSONHandler(w, http.StatusBadRequest, forms.HTTPError{Error: "Expression is empty"})
	default:
		ErrorJSONHandler(w, http.StatusInternalServerError, forms.HTTPError{Error: "Unknown error"})
	}

	JSON(w, forms.Result{Result: result})
}
