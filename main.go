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
	r.Post("/location", h.LocationHandler)
	r.Post("/question", h.CreateQuestion)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(200) })
	log.Printf("Listening on http://0.0.0.0:8080")
	err = http.ListenAndServe("localhost:8080", r)
	if err != nil {
		log.Println(err)
	}
}
