package models

import (
	"time"
)

type Game struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Board       []string   `json:"board"`
	CurrentTurn string     `json:"current_turn"`
	Status      string     `json:"status"`
	PlayerXID   *int       `json:"player_x_id"`
	PlayerOID   *int       `json:"player_o_id"`
	Winner      *string    `json:"winner"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	FinishedAt  *time.Time `json:"finished_at"`
}

type GameRequestBody struct {
	Name   string `json:"name"`
	Role   string `json:"role"`
	UserID int    `json:"user_id"`
}
