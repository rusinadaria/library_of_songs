package service

import (
	"library_of_songs/models"
	"library_of_songs/repository"
)

type SongRepo interface {
	CreateSong(song models.Song) error
    GetAll(filter models.Song, last string, limit string) ([]*models.Song, error)
	GetVerse(id string, limit int, offset int) ([]models.Verse, error)
    // GetSongById(id string) (models.Song, error)
    UpdateSong(id string, song models.Song) error
    DeleteSong(id string) error
	// findSong(song models.Song) models.Song
}

type Service struct {
	SongRepo
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		SongRepo:  NewSongService(repos.SongRepository),
	}
}