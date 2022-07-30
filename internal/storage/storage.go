package storage

import (
	"homework/internal/storage/connections"
	"sync"
	"time"
)

const (
	poolSize        = 10
	timeoutDuration = 3 * time.Second
)

type Storage struct {
	data   map[uint64]*Movie
	lastId uint64

	// mu is used for secure asynchronous access
	mu   sync.RWMutex
	pool *connections.Connections
}

func NewStorage() *Storage {
	var s = &Storage{
		data:   make(map[uint64]*Movie),
		lastId: 1,
		pool:   connections.NewPool(poolSize, timeoutDuration),
	}

	_, _ = s.Add("The Shawshank Redemption", 1994)
	return s
}

func (s *Storage) List() ([]*Movie, error) {
	if err := s.pool.Connect(); err != nil {
		return nil, err
	}

	s.mu.RLock()
	defer func() {
		s.mu.RUnlock()
		s.pool.Disconnect()
	}()

	time.Sleep(time.Second * 10)
	res := make([]*Movie, 0, len(s.data))
	for _, v := range s.data {
		res = append(res, v)
	}
	return res, nil
}

func (s *Storage) Add(title string, year int) (*Movie, error) {
	if err := s.pool.Connect(); err != nil {
		return nil, err
	}

	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		s.pool.Disconnect()
	}()

	m, err := newMovie(title, year, s.lastId)
	if err != nil {
		return nil, err
	}
	if _, ok := s.data[m.Id()]; ok {
		return nil, ErrMovieExists
	}
	s.lastId++
	s.data[m.Id()] = m
	return m, nil
}

func (s *Storage) Update(id uint64, title string, year int) (*Movie, error) {
	if err := s.pool.Connect(); err != nil {
		return nil, err
	}

	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		s.pool.Disconnect()
	}()

	if _, ok := s.data[id]; !ok {
		return nil, ErrMovieNotExists
	}
	m := s.data[id]
	if err := m.SetTitle(title); err != nil {
		return nil, err
	}
	if year != 0 {
		if err := m.SetYear(year); err != nil {
			return nil, err
		}
	}

	return m, nil
}

func (s *Storage) Delete(id uint64) error {
	if err := s.pool.Connect(); err != nil {
		return err
	}

	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		s.pool.Disconnect()
	}()

	if _, ok := s.data[id]; !ok {
		return ErrMovieNotExists
	}
	delete(s.data, id)
	return nil
}
