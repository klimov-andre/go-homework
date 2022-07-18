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
	listCommand = "list"
	addCommand  = "add"
	helpCommand = "help"
)

func listFunc(_ string) string {
	data := storage.List()
	res := make([]string, 0, len(data))
	for _, v := range data {
		res = append(res, v.String())
	}
	return strings.Join(res, "\n")
}

func helpFunc(_ string) string {
	return "/help - list commands\n" +
		"/list - list movies\n" +
		"/add <title> <year> - add new movie with title and year"
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
