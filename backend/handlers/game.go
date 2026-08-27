package handlers

import (
	"backend/models"
	"encoding/json"
	"log"
	"net/http"
)

func (h *Handler) CreateGame(w http.ResponseWriter, r *http.Request) {

	var gameBody models.GameRequestBody
	if err := json.NewDecoder(r.Body).Decode(&gameBody); err != nil {
		http.Error(w, "Invalid json data: "+err.Error(), http.StatusBadRequest)
		return
	}

	if gameBody.Role != "x" && gameBody.Role != "o" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	userID, ok := r.Context().Value("user_id").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	gameBody.UserID = userID
	id, err := h.service.CreateGame(r.Context(), gameBody)
	if err != nil {
		log.Printf("CreateGame error: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func (h *Handler) JoinGame(w http.ResponseWriter, r *http.Request) {

	roomId := r.PathValue("id")

	userID, ok := r.Context().Value("user_id").(int)
	if !ok {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	role, err := h.service.JoinGame(r.Context(), userID, roomId)
	if err != nil {
		log.Printf("JOIN ROOM error: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	log.Printf("JOINED ROOM")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"role": role})
}
