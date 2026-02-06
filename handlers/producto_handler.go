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

func (h *ProductoHandler) Create(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)

	var p models.Producto
	json.NewDecoder(r.Body).Decode(&p)

	if err := h.service.Create(&p, claims.UserID); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Producto creado",
	})
}

func (h *ProductoHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)

	productos, _ := h.service.GetAll(claims.UserID)
	json.NewEncoder(w).Encode(productos)
}
func (h *ProductoHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	p, err := h.service.GetByID(id, claims.UserID)
	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	json.NewEncoder(w).Encode(p)
}
func (h *ProductoHandler) Update(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	var p models.Producto
	json.NewDecoder(r.Body).Decode(&p)

	if err := h.service.Update(id, &p, claims.UserID); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Producto actualizado",
	})
}
func (h *ProductoHandler) Delete(w http.ResponseWriter, r *http.Request) {
	claims := r.Context().Value(middleware.UserContextKey).(*middleware.Claims)
	id, _ := strconv.Atoi(mux.Vars(r)["id"])

	if err := h.service.Delete(id, claims.UserID); err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"message": "Producto eliminado",
	})
}
