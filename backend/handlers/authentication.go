package handlers

import (
	"backend/models"
	"backend/storage"
	"backend/utils"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	service *storage.Storage
}

func NewHandler(service *storage.Storage) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Register(w http.ResponseWriter, r *http.Request) {

	var userBody models.UserBody

	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		http.Error(w, "Invalid Response", http.StatusBadRequest)
		return
	}

	id, _ := h.service.Registred(r.Context(), userBody)
	if id != -1 {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	id, err := h.service.Create(r.Context(), userBody)
	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	log.Printf("Succesfully created user by id %d", id)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]int{"id": id})
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {

	var userBody models.UserBody
	if err := json.NewDecoder(r.Body).Decode(&userBody); err != nil {
		http.Error(w, "Invalid Response", http.StatusBadRequest)
		return
	}
	id, err := h.service.Registred(r.Context(), userBody)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	if id == -1 {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}
	token, err := utils.GenerateJWT(id, userBody.Username)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   24 * 60 * 60,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"redirect": "/home",
		"message":  "Login successful",
	})
}
