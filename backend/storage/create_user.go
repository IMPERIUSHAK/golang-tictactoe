package storage

import (
	"backend/models"
	"context"
	"errors"
)

func (s *Storage) Create(ctx context.Context, user models.UserBody) (int, error) {

	var id int

	err := s.pool.QueryRow(
		ctx,
		"INSERT INTO users (username, password_hash) VALUES($1, $2) RETURNING id",
		user.Username,
		user.Password,
	).Scan(&id)

	return id, err
}

func (s *Storage) Registred(ctx context.Context, user models.UserBody) (int, error) {
	var id int

	err := s.pool.QueryRow(
		ctx,
		`SELECT id FROM users WHERE username = $1 AND password_hash = $2`,
		user.Username,
		user.Password,
	).Scan(&id)

	if err != nil {
		return -1, err
	}

	return id, nil
}

func (s *Storage) Delete(ctx context.Context, id int) error {

	result, err := s.pool.Exec(ctx,
		"DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("Not Found")
	}

	return nil
}
