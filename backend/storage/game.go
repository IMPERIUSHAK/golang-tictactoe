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

func (s *Storage) JoinGame(ctx context.Context, userId int, gameID string) (string, error) {
	var role string
	err := s.pool.QueryRow(
		ctx,
		`UPDATE games
		 SET player_o_id = CASE WHEN player_x_id IS NOT NULL AND player_o_id IS NULL THEN $1 ELSE player_o_id END,
		     player_x_id = CASE WHEN player_x_id IS NULL THEN $1 ELSE player_x_id END
		 WHERE id = $2
		   AND (player_x_id IS NULL OR player_o_id IS NULL)
		   AND player_x_id IS DISTINCT FROM $1
		   AND player_o_id IS DISTINCT FROM $1
		 RETURNING CASE WHEN player_x_id = $1 THEN 'x' ELSE 'o' END`,
		userId,
		gameID,
	).Scan(&role)

	if err != nil {
		return "", err
	}
	return role, nil
}
