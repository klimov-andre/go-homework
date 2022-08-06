package local

import (
	storagePkg "homework/internal/storage"
	"homework/internal/storage/connections"
	"homework/internal/storage/models"
	"sync"
	"time"
)

const (
	poolSize        = 10
	timeoutDuration = 3 * time.Second
)

var _ storagePkg.Storage = (*storage)(nil)

type storage struct {
	data   map[uint64]*models.Movie
	lastId uint64

	// mu is used for secure asynchronous access
	mu   sync.RWMutex
	pool *connections.Connections
}

func NewLocalStorage() storagePkg.Storage {
	var s = &storage{
		data:   make(map[uint64]*models.Movie),
		lastId: 1,
		pool:   connections.NewPool(poolSize, timeoutDuration),
	}

	m, _ := models.NewMovie("The Shawshank Redemption", 1994)
	_ = s.Add(m)
	return s
}

func (s *storage) List(_, _ int) ([]*models.Movie, error) {
	if err := s.pool.Connect(); err != nil {
		return nil, err
	}

	s.mu.RLock()
	defer func() {
		s.mu.RUnlock()
		s.pool.Disconnect()
	}()

	res := make([]*models.Movie, 0, len(s.data))
	for _, v := range s.data {
		res = append(res, v)
	}
	return res, nil
}

func (s *storage) Add(m *models.Movie) error {
	if err := s.pool.Connect(); err != nil {
		return err
	}

	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		s.pool.Disconnect()
	}()

	m.SetId(s.lastId)

	if _, ok := s.data[m.Id]; ok {
		return storagePkg.ErrMovieExists
	}
	s.lastId++
	s.data[m.Id] = m
	return nil
}

func (s *storage) Update(id uint64, newMovie *models.Movie) (*models.Movie, error) {
	if err := s.pool.Connect(); err != nil {
		return nil, err
	}

	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		s.pool.Disconnect()
	}()

	if _, ok := s.data[id]; !ok {
		return nil, storagePkg.ErrMovieNotExists
	}
	m := s.data[id]
	if err := m.SetTitle(newMovie.Title); err != nil {
		return nil, err
	}
	if newMovie.Year != 0 {
		if err := m.SetYear(newMovie.Year); err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (s *storage) Delete(id uint64) error {
	if err := s.pool.Connect(); err != nil {
		return err
	}

	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		s.pool.Disconnect()
	}()

	if _, ok := s.data[id]; !ok {
		return storagePkg.ErrMovieNotExists
	}
	delete(s.data, id)
	return nil
}
