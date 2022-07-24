package storage

import (
	"fmt"
	"github.com/pkg/errors"
)

type Movie struct {
	id    uint64
	title string
	year  int
}

func newMovie(title string, year int, id uint64) (*Movie, error) {
	m := Movie{}
	if err := m.SetTitle(title); err != nil {
		return nil, err
	}
	if err := m.SetYear(year); err != nil {
		return nil, err
	}
	m.SetId(id)
	return &m, nil
}

func (m Movie) Id() uint64 {
	return m.id
}

func (m *Movie) SetId(id uint64) {
	m.id = id
}

func (m Movie) Title() string {
	return m.title
}

func (m *Movie) SetTitle(title string) error {
	if len(title) > 250 {
		return errors.Wrap(TooLongTitle, title)
	}
	m.title = title
	return nil
}

func (m Movie) Year() int {
	return m.year
}

func (m *Movie) SetYear(year int) error {
	if year < 1985 {
		return errors.Wrap(BadYear, fmt.Sprintf("%v", year))
	}
	m.year = year
	return nil
}

func (m Movie) String() string {
	return fmt.Sprintf("%d: %s / %d", m.Id(), m.Title(), m.Year())
}