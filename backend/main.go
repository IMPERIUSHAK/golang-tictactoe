package main

import (
	"backend/db"
	"backend/handlers"
	"backend/middleware"
	"backend/storage"
	"log"
	"net/http"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db.RunMigrations()

	pool, err := db.Connect()
	if err != nil {
		log.Fatal("Error while trying to connect to db")
	}
	defer pool.Close()

	storage := storage.NewStorage(pool)
	handler := handlers.NewHandler(storage)
	mux := http.NewServeMux()
	port := ":9090"

	fileServer := http.FileServer(http.Dir("../static"))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("GET /", handler.LoginPage)
	mux.HandleFunc("POST /api/register", handler.Register)
	mux.HandleFunc("POST /api/login", handler.Login)

	protected := http.NewServeMux()
	protected.HandleFunc("GET /home", handler.HomePage)
	protected.HandleFunc("POST /api/game/new", handler.CreateGame)
	protected.HandleFunc("GET /api/game/join/{id}", handler.JoinGame)

	mux.Handle("GET /home", middleware.AuthMiddleware(protected))
	mux.Handle("GET /api/game/", middleware.AuthMiddleware(protected))
	mux.Handle("POST /api/game/", middleware.AuthMiddleware(protected))

	middle := middleware.Logs(mux)

	log.Printf("Server starting on %s", port)
	if err := http.ListenAndServe(port, middle); err != nil {
		log.Fatal("Error while trying to serve: ", err)
	}
}
