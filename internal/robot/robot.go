package robot

import (
	"github.com/pkg/errors"
	"homework/internal/storage"
	"log"
	"strconv"
	"strings"
)

type Robot struct {
	storage *storage.Storage
}

func NewRobot(storage *storage.Storage) (*Robot, error) {
	return &Robot{storage: storage}, nil
}

func (r *Robot) List() ([]*storage.Movie, error) {
	data, err := r.storage.List()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, EmptyList
	}
	return data, nil
}

func (r *Robot) HelpFunc() string {
	return "/help - list commands\n" +
		"/list - list movies\n" +
		"/add <title> <year> - add new movie with title and year\n" +
		"/remove <id> - remove movie with id\n" +
		"/update <id> <title> {<year>} - remove movie with id and title, year is optional"
}

func (r *Robot) Add(args string) (*storage.Movie, error) {
	log.Printf("add command param: <%s>", args)
	params := strings.Split(args, " ")
	if len(params) != 2 {
		return nil, errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params)
	}

	year, err := strconv.Atoi(params[1])
	if err != nil {
		return nil, errors.Wrapf(BadArgument, "%s", params[1])
	}

	m, err := storage.NewMovie(params[0], year)
	if err != nil {
		return nil, err
	}

	return m, r.storage.Add(m)
}

func (r *Robot) Remove(args string) error {
	log.Printf("remove command param: <%s>", args)
	params := strings.Split(args, " ")
	if len(params) != 1 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params)
	}

	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return errors.Wrapf(BadArgument, "%s", params[0])
	}

	if err = r.storage.Delete(id); err != nil {
		return err
	}

	return nil
}

func (r *Robot) Update(args string) (*storage.Movie, error) {
	log.Printf("update command param: <%s>", args)
	params := strings.Split(args, " ")
	if len(params) < 2 || len(params) > 3 {
		return nil, errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params)
	}

	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return nil, errors.Wrapf(BadArgument, "%s", params[0])
	}

	year := 0
	if len(params) == 3 {
		year, err = strconv.Atoi(params[2])
		if err != nil {
			return nil, errors.Wrapf(BadArgument, "%s", params[0])
		}
	}

	m, err := storage.NewMovie(params[1], year)
	if err != nil {
		return nil, err
	}

	return r.storage.Update(id, m)
}
