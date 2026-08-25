package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type Game struct {
	ID          string    `json:"gameId"`
	Board       [9]string `json:"board"`
	CurrentTurn string    `json:"currentTurn"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
}

var games = make(map[string]*Game)

func CreateGame(w http.ResponseWriter, r *http.Request) {

	var gameId string
	if err := json.NewDecoder(r.Body).Decode(&gameId); err != nil {
		http.Error(w, "Invalid json data", http.StatusBadRequest)
		return
	}

	game := &Game{
		ID:          gameId,
		Board:       [9]string{"", "", "", "", "", "", "", "", ""},
		CurrentTurn: "X",
		Status:      "waiting",
		CreatedAt:   time.Now(),
	}

	games[game.ID] = game

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(game)
}

func JoinGame(w http.ResponseWriter, r *http.Request) {

	gameId := r.PathValue("id")
	game, err := games[gameId]
	if !err {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(game)
}
