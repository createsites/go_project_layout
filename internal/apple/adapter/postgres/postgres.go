package postgres

import (
	"github.com/golang-school/layout/pkg/postgres"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	pool *pgxpool.Pool
}

func New(p *postgres.Postgres) *Postgres {
	return &Postgres{
		pool: p.Pool(),
	}
}
