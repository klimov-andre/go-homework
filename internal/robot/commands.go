package robot

import (
	"fmt"
	"github.com/pkg/errors"
	"homework/internal/storage"
	"log"
	"strconv"
	"strings"
)

var (
	helpCommand = "help"
	listCommand = "list"
	addCommand  = "add"
	removeCommand = "remove"
	updateCommand = "update"
)

func listFunc(_ string) string {
	data := storage.List()
	if len(data) == 0 {
		return EmptyList.Error()
	}
	res := make([]string, 0, len(data))
	for _, v := range data {
		res = append(res, v.String())
	}
	return strings.Join(res, "\n")
}

func helpFunc(_ string) string {
	return "/help - list commands\n" +
		"/list - list movies\n" +
		"/add <title> <year> - add new movie with title and year\n" +
		"/remove <id> - remove movie with id\n" +
		"/update <id> <title> {<year>} - remove movie with id and title, year is optional"
}

func addFunc(args string) string {
	log.Printf("add command param: <%s>", args)
	params := strings.Split(args, " ")
	if len(params) != 2 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}

	year, err := strconv.Atoi(params[1])
	if err != nil {
		return errors.Wrapf(BadArgument, "%s", params[1]).Error()
	}

	m, err := storage.NewMovie(params[0], year)
	if err != nil {
		return err.Error()
	}

	err = storage.Add(m)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("movie %v added", m)
}

func removeFunc(args string) string {
	log.Printf("add command param: <%s>", args)
	params := strings.Split(args, " ")
	if len(params) != 1 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}

	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return errors.Wrapf(BadArgument, "%s", params[0]).Error()
	}

	if err = storage.Delete(id); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("movie %v deleted", id)
}

func updateFunc(args string) string {
	log.Printf("update command param: <%s>", args)
	params := strings.Split(args, " ")
	if len(params) < 2 || len(params) > 3 {
		return errors.Wrapf(BadArgument, "%d items: <%v>", len(params), params).Error()
	}

	id, err := strconv.ParseUint(params[0], 10, 64)
	if err != nil {
		return errors.Wrapf(BadArgument, "%s", params[0]).Error()
	}

	year := 0
	if len(params) == 3 {
		year, err = strconv.Atoi(params[2])
		if err != nil {
			return errors.Wrapf(BadArgument, "%s", params[0]).Error()
		}
	}

	if err = storage.Update(id, params[1], year); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("movie %v successfully updated", id)
}