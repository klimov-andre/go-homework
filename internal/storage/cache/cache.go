package cache

import (
	"context"
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	"sync"
	"time"
)

const (
	poolSize        = 10
	timeoutDuration = 15 * time.Second
)

func NewCache() CacheInterface {
	var s = &Cache{
		data:   make(map[uint64]*models.Movie),
		pool:   connections.NewPool(poolSize, timeoutDuration),
	}

	return s
}

//go:generate mockery --name=CacheInterface --case=snake --with-expecter --structname=CacheInterface --exported=true

type CacheInterface interface{
	AddOrUpdate(ctx context.Context, id uint64, m *models.Movie)
	Delete(ctx context.Context, id uint64) error
	GetById(ctx context.Context, id uint64) (*models.Movie, error)
}

// Cache implements in-memory movie cache for storage
type Cache struct {
	data   map[uint64]*models.Movie

	// mu is used for secure asynchronous access
	mu   sync.RWMutex
	pool *connections.Connections
}

// AddOrUpdate adds movie with specified id to cache,
// if movie with specified id already in cache movie data will be updated
func (s *Cache) AddOrUpdate(_ context.Context, id uint64, m *models.Movie) {
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
	}()

	if m.Id == 0 {
		m.Id = id
	}
	s.data[id] = m
}

// Delete removes movie with specified id from cache
// if movie is not exists return ErrCacheNotExists error
func (s *Cache) Delete(_ context.Context, id uint64) error {
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
	}()

	if _, ok := s.data[id]; !ok {
		return ErrCacheNotExists
	}
	delete(s.data, id)
	return nil
}

// GetById returns movie by id from cache,
// if movie is not exists return ErrCacheNotExists error
func (s *Cache) GetById(_ context.Context, id uint64) (*models.Movie, error) {
	if err := s.pool.Connect(); err != nil {
		return nil, err
	}

	s.mu.RLock()
	defer func() {
		s.mu.RUnlock()
		s.pool.Disconnect()
	}()

	if _, ok := s.data[id]; !ok {
		return nil, ErrCacheNotExists
	}
	return s.data[id], nil
}

