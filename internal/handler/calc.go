package handler

import (
	"encoding/json"
	"net/http"

	"github.com/Irurnnen/ordinary-calc/internal/forms"
	"github.com/Irurnnen/ordinary-calc/internal/models"
	"github.com/Irurnnen/ordinary-calc/pkg/calc"
)

// CalcHandler godoc
//
//	@Summary		Calculate expression
//	@Description	get answer by expression
//	@Tags			Calculator
//	@Param			Expression	body	forms.Expression	true	"Expression"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	models.Result
//	@Success		400	{object}	forms.HTTPError
//	@Failure		422	{object}	forms.HTTPError
//	@Failure		500	{object}	forms.HTTPError
//	@Router			/calculate [post]
func CalcHandler(w http.ResponseWriter, r *http.Request) {
	// Get data from request
	var expression forms.Expression

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
		ErrorJSONHandler(w, http.StatusUnprocessableEntity, forms.HTTPError{Error: "Expression has unpaired brackets"})
		return
	case calc.ErrWrongBracketOrder:
		ErrorJSONHandler(w, http.StatusUnprocessableEntity, forms.HTTPError{Error: "Expression has wrong bracket order"})
		return
	case calc.ErrMultipleOperands:
		ErrorJSONHandler(w, http.StatusUnprocessableEntity, forms.HTTPError{Error: "Expression has multiple operands"})
		return
	case calc.ErrMultipleNumbers:
		ErrorJSONHandler(w, http.StatusUnprocessableEntity, forms.HTTPError{Error: "Expression has multiple sequential numbers"})
		return
	case calc.ErrZeroByDivision:
		ErrorJSONHandler(w, http.StatusUnprocessableEntity, forms.HTTPError{Error: "Expression has zero by division"})
		return
	case calc.ErrExtraOperands:
		ErrorJSONHandler(w, http.StatusUnprocessableEntity, forms.HTTPError{Error: "Expression has at the beginning or at the end"})
		return
	case calc.ErrEmptyExpression:
		ErrorJSONHandler(w, http.StatusUnprocessableEntity, forms.HTTPError{Error: "Expression is empty"})
	default:
		ErrorJSONHandler(w, http.StatusInternalServerError, forms.HTTPError{Error: "Unknown error"})
	}

	JSON(w, models.Result{Result: result})
}
