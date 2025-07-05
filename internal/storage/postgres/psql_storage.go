package postgres

import (
	"bot-templates-profi/internal/config"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	cfg *config.PostgresConfig
	*sqlx.DB
}

func New(cfg *config.PostgresConfig) *Postgres {
	return &Postgres{cfg: cfg}
}

func (p *Postgres) Connect() error {
	cfg := p.cfg
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&client_encoding=%s",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Name,
		"UTF8",
	)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		return err
	}
	p.DB = db

	if err := p.Ping(); err != nil {
		return err
	}

	return nil
}

func (p *Postgres) Close() error {
	return p.DB.Close()
}

func (p *Postgres) Ping() error {
	return p.DB.Ping()
}
