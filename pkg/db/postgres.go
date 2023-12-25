package db

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
)

type DBConfig struct {
	DBURL string
}

func NewPostgres(cfg DBConfig) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalln(err)
	}

	if err != nil {
		return nil, fmt.Errorf("initialize db: %v", err)
	}

	return db, err
}
