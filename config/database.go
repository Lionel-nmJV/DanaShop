package config

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"log"
	"time"
)

func NewDB(config DBConfig) (*sqlx.DB, error) {
	DB_URI := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Pass, config.Name)

	db, err := sqlx.Open("postgres", DB_URI)
	if err != nil {
		return nil, err
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(60 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	log.Println("Database Connected")

	return db, nil
}
