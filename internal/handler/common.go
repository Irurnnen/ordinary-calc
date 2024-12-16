package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ErrorJSONHandler(w http.ResponseWriter, errorCode int, jsonObj any) {
	w.WriteHeader(errorCode)

	JSON(w, jsonObj)
}

func JSON(w http.ResponseWriter, jsonObj any) {
	h := w.Header()

	// Delete the Content-Length header in order to be sure that the
	// entire message reaches the user
	h.Del("Content-Length")

	// Set type of response is JSON
	h.Set("Content-Type", "application/json; charset=utf-8")

	err := json.NewEncoder(w).Encode(jsonObj)
	if err != nil {
		fmt.Fprint(w, "\nError while encode to json\n")
	}
}
