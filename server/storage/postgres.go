package storage

import "github.com/jmoiron/sqlx"

func NewPostgresDb() (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", "postgres://user:password@db:5432/distr?sslmode=disable")
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
