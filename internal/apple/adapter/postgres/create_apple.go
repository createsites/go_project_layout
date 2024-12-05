package postgres

import (
	"context"
	"fmt"
	"github.com/doug-martin/goqu/v9"
	"github.com/golang-school/layout/internal/apple/entity/apple"
	"github.com/golang-school/layout/pkg/tracer"
)

func (p *Postgres) CreateApple(ctx context.Context, a apple.Apple) (err error) {
	ctx, span := tracer.Start(ctx, "postgres CreateApple")
	defer tracer.End(span)

	dataset := goqu.Insert("apple").
		Rows(goqu.Record{
			"id":     a.ID,
			"name":   a.Name,
			"status": a.Status,
		})

	sql, _, err := dataset.ToSQL()
	if err != nil {
		return fmt.Errorf("dataset.ToSQL: %w", err)
	}

	_, err = p.pool.Exec(ctx, sql)
	if err != nil {
		return fmt.Errorf("r.pool.Exec: %w", err)
	}

	return nil
}