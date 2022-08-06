package storage

import (
	"homework/internal/storage/models"
)

type Storage interface {
	List(limit, offset int) ([]*models.Movie, error)
	Add(m *models.Movie) error
	Update(id uint64, newMovie *models.Movie) (*models.Movie, error)
	Delete(id uint64) error
}