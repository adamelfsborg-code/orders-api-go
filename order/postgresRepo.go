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
	return p.Client.Ping(ctx)
}

func (p *PostgresRepo) Close() error {
	return p.Client.Close()
}

func (p *PostgresRepo) Create(ctx context.Context, order Order) error {
	_, err := p.Client.Model(&order).Insert()
	if err != nil {
		return fmt.Errorf("Failed to create order: %w", err)
	}
	return nil
}

func (p *PostgresRepo) List(ctx context.Context, page FindAllPage) (FindResult, error) {
	return FindResult{}, nil
}

func (p *PostgresRepo) FindByID(ctx context.Context, id uint64) (Order, error) {
	return Order{}, nil
}

func (p *PostgresRepo) UpdateByID(ctx context.Context, order Order) error {
	return nil
}

func (p *PostgresRepo) DeleteByID(ctx context.Context, id uint64) error {
	return nil
}
