package local

import (
	"context"
	cachePkg "homework/internal/storage/cache"
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	"sync"
	"time"
)

const (
	poolSize        = 10
	timeoutDuration = 15 * time.Second
)

func NewLocalCache() cachePkg.CacheInterface {
	var s = &cache{
		data:   make(map[uint64]*models.Movie),
		pool:   connections.NewPool(poolSize, timeoutDuration),
	}

	return s
}

// cache implements in-memory movie cache for storage
type cache struct {
	data   map[uint64]*models.Movie

	// mu is used for secure asynchronous access
	mu   sync.RWMutex
	pool *connections.Connections
}

// AddOrUpdate adds movie with specified id to cache,
// if movie with specified id already in cache movie data will be updated
func (c *cache) AddOrUpdate(_ context.Context, id uint64, m *models.Movie) {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()

	if m.Id == 0 {
		m.Id = id
	}
	c.data[id] = m
}

// Delete removes movie with specified id from cache
// if movie is not exists return ErrCacheNotExists error
func (c *cache) Delete(_ context.Context, id uint64) error {
	c.mu.Lock()
	defer func() {
		c.mu.Unlock()
	}()

	if _, ok := c.data[id]; !ok {
		return cachePkg.ErrCacheNotExists
	}
	delete(c.data, id)
	return nil
}

// GetById returns movie by id from cache,
// if movie is not exists return ErrCacheNotExists error
func (c *cache) GetById(_ context.Context, id uint64) (*models.Movie, error) {
	if err := c.pool.Connect(); err != nil {
		return nil, err
	}

	c.mu.RLock()
	defer func() {
		c.mu.RUnlock()
		c.pool.Disconnect()
	}()

	if _, ok := c.data[id]; !ok {
		return nil, cachePkg.ErrCacheNotExists
	}
	return c.data[id], nil
}

