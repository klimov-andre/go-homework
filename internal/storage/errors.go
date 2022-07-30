package storage

import "github.com/pkg/errors"

var (
	ErrValidation     = errors.New("invalid data")
	ErrMovieExists    = errors.New("movie already exists")
	ErrMovieNotExists = errors.New("movie is not exists")
)
