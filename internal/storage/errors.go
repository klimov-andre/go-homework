package storage

import "github.com/pkg/errors"

var (
	ErrMovieExists    = errors.New("movie already exists")
	ErrMovieNotExists = errors.New("movie is not exists")
)
