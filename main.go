package main

import (
	"auth_service/database"
	"auth_service/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Ошибка загрузки .env файла:", err)
	}
	db, err := database.InitDatabase()
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer db.Close()

	router := mux.NewRouter()
	router.HandleFunc("/auth/tokens", handlers.GenerateTokens(db)).Methods("POST")
	router.HandleFunc("/auth/refresh", handlers.RefreshTokens(db)).Methods("POST")

	log.Println("Сервер запущен на порту :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
