package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"time"
)

func NewDB() *sqlx.DB {
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	DB_USER := os.Getenv("DB_USER")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASS := os.Getenv("DB_PASSWORD")

	DB_URI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", DB_USER, DB_PASS, DB_HOST, DB_PORT, DB_NAME)

	db, err := sqlx.Open("postgres", DB_URI)
	if err != nil {
		// Handle error
	}
	defer db.Close()

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	log.Println("Database Connected")

	return db
}
