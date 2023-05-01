package postgres

import (
	"context"

	"github.com/go-pg/pg/v10"
)

func NewPostgres() (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		User:     "postgres",
		Password: "king1337",
		Addr:     "localhost:5432",
		Database: "coursch",
	})

	err := db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return db, nil
}
