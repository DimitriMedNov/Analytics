package location

import (
	"encoding/json"
	"net/http"

	"github.com/jackc/pgx/v5/pgxpool"
)

type locationInput struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Handler struct{ Pool *pgxpool.Pool }

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
