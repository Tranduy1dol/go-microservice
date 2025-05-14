package store

import (
	"context"
	"database/sql"
)

type UsersStorage struct {
	db *sql.DB
}

func (s *UsersStorage) Create(ctx context.Context) error {
	// Implement the logic to create a user in the database
	return nil
}
