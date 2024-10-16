package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v4"
)

func init() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
}

func main() {
    conn, err := pgx.Connect(context.Background(), os.Getenv("POSTGRES_URL"))
    if err != nil {
        log.Fatalf("Unable to connect to database: %v\n", err)
    }
    defer conn.Close(context.Background())

    fmt.Println("Connected to PostgreSQL!")
}
