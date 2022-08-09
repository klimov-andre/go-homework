package facade

import (
	"context"
	"github.com/pkg/errors"
	"homework/internal/storage/cache"
	"homework/internal/storage/db"
	"homework/internal/storage/models"
)

var _ StorageFacade = (*storageFacade)(nil)

type StorageFacade interface {
	List(ctx context.Context, limit, offset int) ([]*models.Movie, error)
	Add(ctx context.Context, m *models.Movie) error
	Update(ctx context.Context, id uint64, newMovie *models.Movie) (*models.Movie, error)
	Delete(ctx context.Context, id uint64) error
	GetOneMovie(ctx context.Context, id uint64) (*models.Movie, error)
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

func (s *storageFacade) List(ctx context.Context, limit, offset int) ([]*models.Movie, error) {
	movies, err := s.db.List(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	for _, m := range movies {
		s.cache.AddOrUpdate(ctx, m.Id, m)
	}

	return movies, nil
}

func (s *storageFacade) Add(ctx context.Context, m *models.Movie) error {
	if err := s.db.Add(ctx, m); err != nil {
		return err
	}

	s.cache.AddOrUpdate(ctx, m.Id, m)
	return nil
}

func (s *storageFacade) Update(ctx context.Context, id uint64, newMovie *models.Movie) (*models.Movie, error) {
	if _, err := s.db.Update(ctx, id, newMovie); err != nil {
		return nil, err
	}

	s.cache.AddOrUpdate(ctx, id, newMovie)
	return newMovie, nil
}

func (s *storageFacade) Delete(ctx context.Context, id uint64) error {
	if err := s.db.Delete(ctx, id); err != nil {
		return err
	}

	if err := s.cache.Delete(ctx, id); err != nil {
		if !errors.Is(err, cache.ErrCacheNotExists) {
			return err
		}
	}
	return nil
}

func (s *storageFacade) GetOneMovie(ctx context.Context, id uint64) (*models.Movie, error) {
	m, err := s.cache.GetById(ctx, id)
	if err != nil && !errors.Is(err, cache.ErrCacheNotExists) {
		return nil, err
	}
	if m != nil {
		return m, nil
	}

	m, err = s.db.GetOneMovie(ctx, id)
	if err != nil {
		return nil, err
	}
	s.cache.AddOrUpdate(ctx, id, m)

	return m, nil
}