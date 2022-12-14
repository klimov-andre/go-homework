package router

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

type HandlerFunc func(string) (string, error)

type Router struct {
	routes map[string]HandlerFunc
}

func New() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}

func (r *Router) AddCommand(command string, handlerFunc HandlerFunc) {
	r.routes[command] = handlerFunc
}

func (r *Router) Handle(command, arguments string) (string, error) {
	id := uuid.New().String()
	logrus.Infof("received command uuid: %v, command: %v, arguments %v", id, command, arguments)

	if handler, ok := r.routes[command]; ok {
		return handler(arguments)
	}
	return "", UnknownCommand
}
