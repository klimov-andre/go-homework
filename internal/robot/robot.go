package robot

import (
	"context"
	"github.com/pkg/errors"
	"homework/internal/storage/facade"
	"homework/internal/storage/models"
	"log"
	"strconv"
	"strings"
)

const (
	defaultLimit = 100
	defaultOrder = "ASC"
)

type Robot struct {
	storage facade.StorageFacade
}

func NewRobot(storage facade.StorageFacade) (*Robot, error) {
	return &Robot{storage: storage}, nil
}

func (r *Robot) List() ([]*models.Movie, error) {
	data, err := r.storage.List(context.Background(), defaultLimit, 0, defaultOrder)
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

func (r *Robot) Add(args string) (*models.Movie, error) {
	log.Printf("add command param: <%s>", args)
	params := strings.Split(args, " ")
	if len(params) != 2 {
		return nil, errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params)
	}

	year, err := strconv.Atoi(params[1])
	if err != nil {
		return nil, errors.Wrapf(BadArgument, "%s", params[1])
	}

	m, err := models.NewMovie(params[0], year)
	if err != nil {
		return nil, err
	}

	return m, r.storage.Add(context.Background(), m)
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

	if err = r.storage.Delete(context.Background(), id); err != nil {
		return err
	}

	return nil
}

func (r *Robot) Update(args string) (*models.Movie, error) {
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

	m, err := models.NewMovie(params[1], year)
	if err != nil {
		return nil, err
	}

	return m, r.storage.Update(context.Background(), id, m)
}
