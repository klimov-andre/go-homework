package storage

import (
	"sync"
)

const poolSize = 10

type Storage struct {
	data   map[uint64]*Movie
	lastId uint64

	// mu is used for secure asynchronous access
	mu sync.RWMutex
	// connectionPool limits connections count to storage
	connectionPool chan struct{}
}

func NewStorage() *Storage {
	var s = &Storage{
		data:           make(map[uint64]*Movie),
		lastId:         1,
		connectionPool: make(chan struct{}, poolSize),
	}

	_, _ = s.Add("The Shawshank Redemption", 1994)
	return s
}

func (s *Storage) List() []*Movie {
	s.connectionPool <- struct{}{}
	s.mu.RLock()
	defer func() {
		s.mu.RUnlock()
		<-s.connectionPool
	}()

	res := make([]*Movie, 0, len(s.data))
	for _, v := range s.data {
		res = append(res, v)
	}
	return res
}

func (s *Storage) Add(title string, year int) (*Movie, error) {
	s.connectionPool <- struct{}{}
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		<-s.connectionPool
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
	s.connectionPool <- struct{}{}
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		<-s.connectionPool
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
	s.connectionPool <- struct{}{}
	s.mu.Lock()
	defer func() {
		s.mu.Unlock()
		<-s.connectionPool
	}()

	if _, ok := s.data[id]; !ok {
		return ErrMovieNotExists
	}
	delete(s.data, id)
	return nil
}
