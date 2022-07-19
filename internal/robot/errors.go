package robot

import "github.com/pkg/errors"

var (
	BadArgument = errors.New("bad argument")
	EmptyList = errors.New("empty list")
)
