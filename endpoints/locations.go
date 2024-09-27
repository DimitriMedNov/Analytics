package endpoints

import (
	"analytics/dbgen"
	"context"
	"encoding/json"
	"math/big"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

type locationInput struct {
	Lat  int64 `json:"lat"`
	Long int64 `json:"long"`
}

func (h Handler) LocationHandler(w http.ResponseWriter, r *http.Request) {
	var input locationInput

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	savedValues, err := dbgen.New(h.Pool).CreateLocation(context.Background(), dbgen.CreateLocationParams{
		Lat:  pgtype.Numeric{Int: big.NewInt(input.Lat), Exp: -7},
		Long: pgtype.Numeric{Int: big.NewInt(input.Long), Exp: -7},
	})

	response, err := json.Marshal(savedValues)
	if err != nil {
		http.Error(w, "Error processing input", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
