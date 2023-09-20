package order

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

type PostgresRepo struct {
	Client *pg.DB
}

func (p *PostgresRepo) Ping(ctx context.Context) error {
	err := p.Client.Ping(ctx)
	if err != nil {
		return fmt.Errorf("Failed to connect to postgres: %w", err)
	}
	return nil
}

func (p *PostgresRepo) Close(ctx context.Context) error {
	return nil
}

func (p *PostgresRepo) Create(ctx context.Context, order OrderModel) error {
	return nil
}

func (p *PostgresRepo) List(ctx context.Context, page FindAllPage) (FindResult, error) {
	return FindResult{}, nil
}

func (p *PostgresRepo) FindByID(ctx context.Context, id uint64) (OrderModel, error) {
	return OrderModel{}, nil
}

func (p *PostgresRepo) UpdateByID(ctx context.Context, order OrderModel) error {
	return nil
}

func (p *PostgresRepo) DeleteByID(ctx context.Context, id uint64) error {
	return nil
}
