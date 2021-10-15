package mysql

import (
	"Gymondo/platform/mysql"
	"context"
)

type Repository struct {
	database *mysql.Connection
}

func CreateRepository(db *mysql.Connection) (*Repository, error) {
	return &Repository{
		database: db,
	}, nil
}

func (m *Repository) GetProducts(ctx context.Context, limit int) error {
	return nil
}
