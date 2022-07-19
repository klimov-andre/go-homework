package storage

var data map[uint64]*Movie

func init() {
	data = make(map[uint64]*Movie)

	u, _ := NewMovie("The Shawshank Redemption", 1994)
	data[u.id] = u
}

func List() []*Movie {
	res := make([]*Movie, 0, len(data))
	for _, v := range data {
		res = append(res, v)
	}
	return res
}

func Add(m *Movie) error {
	if _, ok := data[m.Id()]; ok {
		return MovieExists
	}
	data[m.Id()] = m
	return nil
}

func Update(id uint64, title string, year int) error {
	if _, ok := data[id]; !ok {
		return MovieNotExists
	}
	u := data[id]
	if err := u.SetTitle(title); err != nil {
		return err
	}
	if year != 0 {
		if err := u.SetYear(year); err != nil {
			return err
		}
	}

	return nil
}

func Delete(id uint64) error {
	if _, ok := data[id]; !ok {
		return MovieNotExists
	}
	delete(data, id)
	return nil
}
