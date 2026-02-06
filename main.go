package main

import (
	"log"
	"net/http"
	"os"

	"apiven/config"
	"apiven/routes"

	"github.com/joho/godotenv"
)

func main() {

	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No se pudo cargar el archivo .env")
	}

	// Conectar a la BD
	config.ConnectDB()
	defer config.CloseDB()

	// Rutas
	router := routes.SetupRoutes()

	// Puerto del servidor
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("🚀 Servidor corriendo en el puerto", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
