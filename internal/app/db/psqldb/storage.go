package psqldb

import (
	"bot-templates-profi/internal/config"
	"bot-templates-profi/internal/storage/postgres"
)

type PsqlStorage struct {
	cfg *config.PostgresConfig
}

func Run(cfg *config.PostgresConfig) (*postgres.Postgres, error) {
	psql := postgres.New(cfg)

	if err := psql.Connect(); err != nil {
		return nil, err
	}

	return psql, nil
}

func MustRun(cfg *config.PostgresConfig) *postgres.Postgres {
	psql, err := Run(cfg)
	if err != nil {
		panic(err)
	}

	return psql
}
