package main

import (
	"analytics/endpoints"
	"context"
	"embed"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed db/migrations/*.sql
var fs embed.FS

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL no está definida")
	}

	u, err := url.Parse(dbURL)
	if err != nil {
		log.Fatalf("Error al parsear DATABASE_URL: %v", err)
	}

	db := dbmate.New(u)
	db.FS = fs

	fmt.Println("Aplicando migraciones...")
	err = db.CreateAndMigrate()
	if err != nil {
		log.Fatalf("Error aplicando migraciones: %v", err)
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("No se puede conectar a la base de datos: %v", err)
	}
	defer pool.Close()

	r := chi.NewRouter()
	h := endpoints.Handler{Pool: pool}

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Definimos las rutas
	r.Post("/location", h.LocationHandler)
	r.Post("/question", h.CreateQuestion)
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) { w.WriteHeader(http.StatusOK) })

	log.Println("Escuchando en http://0.0.0.0:8000")
	err = http.ListenAndServe("0.0.0.0:8000", r)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
