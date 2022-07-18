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
		return UserExists
	}
	data[m.Id()] = m
	return nil
}

func Update(u *Movie) error {
	if _, ok := data[u.Id()]; ok {
		return UserNotExists
	}
	data[u.Id()] = u
	return nil
}

func Delete(u *Movie) error {
	if _, ok := data[u.Id()]; ok {
		return UserNotExists
	}
	delete(data, u.Id())
	return nil
}
