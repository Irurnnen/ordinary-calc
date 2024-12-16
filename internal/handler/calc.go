package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Irurnnen/ordinary-calc/internal/models"
)

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	// Get data from request
	var expression models.Expression

	err := json.NewDecoder(r.Body).Decode(&expression)
	if err != nil {
		http.Error(w, "Provided data is invalid", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Expression: %s", expression.Expression)
	return
}
