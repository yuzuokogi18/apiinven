package handlers

import (
	"apiven/middleware"
	"apiven/models"
	"apiven/services"
	"encoding/json"
	"net/http"
)

type CategoriaHandler struct {
	service *services.CategoriaService
}

func NewCategoriaHandler() *CategoriaHandler {
	return &CategoriaHandler{
		service: services.NewCategoriaService(),
	}
}

func (h *CategoriaHandler) Create(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)

	var c models.Categoria
	json.NewDecoder(r.Body).Decode(&c)

	h.service.Create(&c, claims.UserID)

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Categoría creada",
	})
}
