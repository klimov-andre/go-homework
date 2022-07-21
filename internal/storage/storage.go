package storage

type Storage struct {
	data map[uint64]*Movie
	lastId uint64
}

func NewStorage() *Storage {
	var s = &Storage{
		data: make(map[uint64]*Movie),
		lastId: 1,
	}

	_, _ = s.Add("The Shawshank Redemption", 1994)
	return s
}

func (s *Storage) List() []*Movie {
	res := make([]*Movie, 0, len(s.data))
	for _, v := range s.data {
		res = append(res, v)
	}
	return res
}

func (s *Storage) Add(title string, year int) (*Movie, error) {
	m, err := newMovie(title, year, s.lastId)
	if err != nil {
		return nil, err
	}
	if _, ok := s.data[m.Id()]; ok {
		return nil, MovieExists
	}
	s.lastId++
	s.data[m.Id()] = m
	return m, nil
}

func (s *Storage) Update(id uint64, title string, year int) error {
	if _, ok := s.data[id]; !ok {
		return MovieNotExists
	}
	m := s.data[id]
	if err := m.SetTitle(title); err != nil {
		return err
	}
	if year != 0 {
		if err := m.SetYear(year); err != nil {
			return err
		}
	}

	return nil
}

func (s *Storage) Delete(id uint64) error {
	if _, ok := s.data[id]; !ok {
		return MovieNotExists
	}
	delete(s.data, id)
	return nil
}
