package main

import (
	"analytics/location"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	r := chi.NewRouter()
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Cant connect to the db")
	}
	h := location.Handler{Pool: pool}
	r.Post("/test", h.LocationHandler)
	http.ListenAndServe("localhost:8000", r)
}
