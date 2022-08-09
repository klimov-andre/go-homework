package cache

import (
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	"sync"
	"time"
)

const (
	poolSize        = 10
	timeoutDuration = 15 * time.Second
)

type Cache struct {
	data   map[uint64]*models.Movie

	// mu is used for secure asynchronous access
	mu   sync.RWMutex
	pool *connections.Connections
}

func NewCache() *Cache {
	var s = &Cache{
		data:   make(map[uint64]*models.Movie),
		pool:   connections.NewPool(poolSize, timeoutDuration),
	}

	return s
}

func (s *Cache) AddOrUpdate(id uint64, m *models.Movie) {
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
	}()

	s.data[id] = m
}

func (s *Cache) Delete(id uint64) error {
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

func (s *Cache) GetById(id uint64) (*models.Movie, error) {
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

