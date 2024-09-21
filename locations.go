package main

import (
	"encoding/json"
	"net/http"
)

type locationInput struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

func (h Handler) LocationHandler(w http.ResponseWriter, r *http.Request) {
	var input locationInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	response, err := json.Marshal(input)
	if err != nil {
		http.Error(w, "Error processing input", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
