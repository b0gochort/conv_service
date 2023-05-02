package postgres

import (
	"context"

	"github.com/go-pg/pg/v10"
)

type PG struct {
	Addr         string
	User         string
	Password     string
	DataBaseName string
}

func NewPostgres(cfg PG) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		User:     cfg.User,
		Password: cfg.Password,
		Addr:     cfg.Addr,
		Database: cfg.DataBaseName,
	})

	err := db.Ping(context.Background())
	if err != nil {
		return nil, err
	}
	return db, nil
}
