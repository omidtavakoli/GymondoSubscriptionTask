package mysql

import (
	"Gymondo/platform/postgres"
	"context"
)

type Repository struct {
	database *postgres.Connection
}

func CreateRepository(db *postgres.Connection) (*Repository, error) {
	return &Repository{
		database: db,
	}, nil
}

func (m *Repository) GetProducts(ctx context.Context, limit int) error {
	return nil
}
