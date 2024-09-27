package main

import (
	"analytics/endpoints"
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	r := chi.NewRouter()
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Cant connect to the db")
	}
	h := endpoints.Handler{Pool: pool}
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Post("/test", h.LocationHandler)
	log.Printf("Listening on http://localhost:8000")
	err = http.ListenAndServe("localhost:8000", r)
	if err != nil {
		log.Println(err)
	}
}
