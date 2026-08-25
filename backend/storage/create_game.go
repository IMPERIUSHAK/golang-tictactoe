package storage

import (
	"backend/models"
	"context"
	"database/sql"
)

func (s *Storage) CreateGame(ctx context.Context, game models.GameRequestBody) (string, error) {
	var sidex, sideo sql.NullInt64
	var id string
	if game.Role == "o" {
		sideo.Int64 = int64(game.UserID)
		sideo.Valid = true
		sidex.Valid = false
	} else {
		sidex.Int64 = int64(game.UserID)
		sidex.Valid = true
		sideo.Valid = false
	}
	err := s.pool.QueryRow(
		ctx,
		"INSERT INTO games ( name, player_x_id, player_o_id) VALUES($1, $2, $3) RETURNING id",
		game.Name,
		sidex,
		sideo,
	).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}
