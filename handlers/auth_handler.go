package handlers

import (
	"apiven/models"
	"apiven/services"
	"encoding/json"
	"net/http"
)

type AuthHandler struct {
	service *services.AuthService
}

func NewAuthHandler() *AuthHandler {
	return &AuthHandler{
		service: services.NewAuthService(),
	}
}

// REGISTER
func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	usuario, token, err := h.service.Register(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Usuario registrado correctamente",
		"token":   token,
		"usuario": usuario,
	})
}

// LOGIN
func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	usuario, token, err := h.service.Login(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Login exitoso",
		"token":   token,
		"usuario": usuario,
	})
}
