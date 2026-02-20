package handlers

import (
	"apiven/middleware"
	"apiven/models"
	"apiven/services"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ProductoHandler struct {
	service *services.ProductoService
}

func NewProductoHandler() *ProductoHandler {
	return &ProductoHandler{
		service: services.NewProductoService(),
	}
}

// 🔥 CREATE
func (h *ProductoHandler) Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)
	if !ok {
		http.Error(w, "usuario no autorizado", http.StatusUnauthorized)
		return
	}

	var p models.Producto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "datos inválidos", http.StatusBadRequest)
		return
	}

	// 🔥 Pasamos el userID al service (IMPORTANTE)
	if err := h.service.Create(&p, claims.UserID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Producto creado",
	})
}

// 🔥 GET ALL
func (h *ProductoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)
	if !ok {
		http.Error(w, "usuario no autorizado", http.StatusUnauthorized)
		return
	}

	productos, err := h.service.GetAll(claims.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(productos)
}

// 🔥 GET BY ID
func (h *ProductoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)
	if !ok {
		http.Error(w, "usuario no autorizado", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	p, err := h.service.GetByID(id, claims.UserID)
	if err != nil {
		http.Error(w, "producto no encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(p)
}

// 🔥 UPDATE
func (h *ProductoHandler) Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)
	if !ok {
		http.Error(w, "usuario no autorizado", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	var p models.Producto
	if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
		http.Error(w, "datos inválidos", http.StatusBadRequest)
		return
	}

	if err := h.service.Update(id, &p, claims.UserID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Producto actualizado",
	})
}

// 🔥 DELETE
func (h *ProductoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	claims, ok := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)
	if !ok {
		http.Error(w, "usuario no autorizado", http.StatusUnauthorized)
		return
	}

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "id inválido", http.StatusBadRequest)
		return
	}

	if err := h.service.Delete(id, claims.UserID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Producto eliminado",
	})
}
