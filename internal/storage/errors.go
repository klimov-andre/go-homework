package storage

import "github.com/pkg/errors"

var (
	BadYear       = errors.New("bad year")
	TooLongTitle  = errors.New("title is too long")
	MovieExists    = errors.New("movie already exists")
	MovieNotExists = errors.New("movie is not exists")
)
