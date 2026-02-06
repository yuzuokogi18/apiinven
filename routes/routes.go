package routes

import (
	"apiven/handlers"
	"apiven/middleware"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api").Subrouter()

	auth := handlers.NewAuthHandler()
	producto := handlers.NewProductoHandler()

	api.HandleFunc("/auth/register", auth.Register).Methods("POST")
	api.HandleFunc("/auth/login", auth.Login).Methods("POST")

	api.HandleFunc("/productos", middleware.AuthMiddleware(producto.Create)).Methods("POST")
	api.HandleFunc("/productos", middleware.AuthMiddleware(producto.GetAll)).Methods("GET")
	api.HandleFunc("/productos/{id}", middleware.AuthMiddleware(producto.GetByID)).Methods("GET")
	api.HandleFunc("/productos/{id}", middleware.AuthMiddleware(producto.Update)).Methods("PUT")
	api.HandleFunc("/productos/{id}", middleware.AuthMiddleware(producto.Delete)).Methods("DELETE")
	categoria := handlers.NewCategoriaHandler()

	api.HandleFunc("/categorias", middleware.AuthMiddleware(categoria.Create)).Methods("POST")

	return r
}
