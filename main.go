package main

import (
	// "fmt"
	"log/slog"
	"net/http"
	"os"
	"github.com/joho/godotenv"
	"library_of_songs/internal/service"
	"library_of_songs/internal/handler"
	"library_of_songs/repository"
)

// @title Songs API
// @version 1.0
// @description This is a sample API for managing songs.
// @host localhost:8080
// @BasePath /songs
func main() {
	logger := configLogger()

	if err := godotenv.Load(); err != nil {
		logger.Error("error loading env variables", slog.String("error", err.Error()))
	}

	db, err := repository.ConnectDatabase(logger)
	if err != nil {
		logger.Error("failed to connect db", slog.String("error", err.Error()))
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	h := handler.NewHandler(services)

	err = http.ListenAndServe(os.Getenv("PORT"), h.InitRoutes())
	if err != nil {
		logger.Error("failed start server")
		panic(err)
	}
}

func configLogger() *slog.Logger {
	var logger *slog.Logger

	f, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
        slog.Error("Unable to open a file for writing")
    }

	opts := &slog.HandlerOptions{
        Level: slog.LevelDebug,
    }

	logger = slog.New(slog.NewJSONHandler(f, opts))
	return logger
}