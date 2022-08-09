package models

import (
	"fmt"
	"github.com/pkg/errors"
)

type Movie struct {
	Id    uint64
	Title string
	Year  int
}

func NewMovie(title string, year int) (*Movie, error) {
	m := Movie{}
	if err := m.SetTitle(title); err != nil {
		return nil, err
	}
	if err := m.SetYear(year); err != nil {
		return nil, err
	}
	return &m, nil
}

func (m *Movie) SetId(id uint64) {
	m.Id = id
}

func (m *Movie) SetTitle(title string) error {
	if len(title) > 250 {
		return errors.Wrapf(ErrValidation, "title is too long %s", title)
	}
	m.Title = title
	return nil
}

func (m *Movie) SetYear(year int) error {
	if year < 1985 {
		return errors.Wrapf(ErrValidation, "bad year %v", year)
	}
	m.Year = year
	return nil
}

func (m Movie) String() string {
	return fmt.Sprintf("%d: %s / %d", m.Id, m.Title, m.Year)
}