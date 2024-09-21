package main

import (
	"analytics/dbgen"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
)

func main() {
	pool, err := pgxpool.New(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal("Cant connect to the db")
	}
	r := gin.Default()
	r.GET("/authors", func(c *gin.Context) {
		authors, err := dbgen.New(pool).ListAuthors(context.Background())
		if err != nil {
			log.Fatal("DB error", err)
		}
		c.JSON(200, gin.H{
			"author": authors,
		})
	})
	r.Run("localhost:8000") // listen and serve on 0.0.0.0:8080
}
