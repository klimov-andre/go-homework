package facade

import (
	"github.com/pkg/errors"
	"homework/internal/storage/cache"
	"homework/internal/storage/db"
	"homework/internal/storage/models"
)

var _ StorageFacade = (*storageFacade)(nil)

type StorageFacade interface {
	List(limit, offset int) ([]*models.Movie, error)
	Add(m *models.Movie) error
	Update(id uint64, newMovie *models.Movie) (*models.Movie, error)
	Delete(id uint64) error
	GetOneMovie(id uint64) (*models.Movie, error)
}

type storageFacade struct {
	db *db.Database
	cache *cache.Cache
}

func NewStorage() (StorageFacade, error) {
	db, err := db.NewDatabase()
	if err != nil {
		return nil, errors.Wrap(err, "could not init database")
	}

	return &storageFacade{
		db: db,
		cache: cache.NewCache(),
	}, err
}

func (s *storageFacade) List(limit, offset int) ([]*models.Movie, error) {
	movies, err := s.db.List(limit, offset)
	if err != nil {
		return nil, err
	}

	for _, m := range movies {
		s.cache.AddOrUpdate(m.Id, m)
	}

	return movies, nil
}

func (s *storageFacade) Add(m *models.Movie) error {
	if err := s.db.Add(m); err != nil {
		return err
	}

	s.cache.AddOrUpdate(m.Id, m)
	return nil
}

func (s *storageFacade) Update(id uint64, newMovie *models.Movie) (*models.Movie, error) {
	if _, err := s.db.Update(id, newMovie); err != nil {
		return nil, err
	}

	s.cache.AddOrUpdate(id, newMovie)
	return newMovie, nil
}

func (s *storageFacade) Delete(id uint64) error {
	if err := s.db.Delete(id); err != nil {
		return err
	}

	if err := s.cache.Delete(id); err != nil {
		if !errors.Is(err, cache.ErrCacheNotExists) {
			return err
		}
	}
	return nil
}

func (s *storageFacade) GetOneMovie(id uint64) (*models.Movie, error) {
	m, err := s.cache.GetById(id)
	if err != nil && !errors.Is(err, cache.ErrCacheNotExists) {
		return nil, err
	}
	if m != nil {
		return m, nil
	}

	m, err = s.db.GetOneMovie(id)
	if err != nil {
		return nil, err
	}
	s.cache.AddOrUpdate(id, m)

	return m, nil
}