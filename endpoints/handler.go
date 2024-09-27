package endpoints

import "github.com/jackc/pgx/v5/pgxpool"

type Handler struct{ Pool *pgxpool.Pool }
