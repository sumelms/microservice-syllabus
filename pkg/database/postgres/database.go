package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	// Adding the postgres driver
	_ "github.com/jackc/pgx"
	"github.com/pkg/errors"

	"github.com/sumelms/microservice-syllabus/pkg/config"
)

func Connect(cfg *config.Database) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Database, cfg.Username, cfg.Password)
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect to the database")
	}

	return db, nil
}
