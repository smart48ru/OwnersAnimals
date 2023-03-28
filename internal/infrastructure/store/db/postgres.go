package db

import (
	"OwnersAnimals/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
)

func NewPostgresConnect(cfg config.DB) (*sqlx.DB, error) {
	log.Debug().Msgf("Configuring Postgres")
	// dbase, err := sqlx.Open("postgres", dsn)
	dbase, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=%t",
		cfg.Host, cfg.Post, cfg.User, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		return nil, err
	}
	err = dbase.Ping()
	if err != nil {
		return nil, err
	}

	if err != nil {
		dbase.Close()
		return nil, err
	}

	return dbase, nil
}
