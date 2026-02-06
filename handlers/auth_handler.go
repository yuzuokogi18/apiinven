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

func (h *AuthHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	json.NewDecoder(r.Body).Decode(&req)

	if err := h.service.Register(&req); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Usuario registrado",
	})
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req models.LoginRequest
	json.NewDecoder(r.Body).Decode(&req)

	token, err := h.service.Login(&req)
	if err != nil {
		http.Error(w, err.Error(), 401)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"token": token,
	})
}
