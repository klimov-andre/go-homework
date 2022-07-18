package router

type HandlerFunc func(string) string

type Router struct {
	routes map[string]HandlerFunc
}

func New() *Router {
	return &Router{
		routes: make(map[string]HandlerFunc),
	}
}

func (r *Router) Add(command string, handlerFunc HandlerFunc) {
	r.routes[command] = handlerFunc
}

func (r *Router) Handle(command, arguments string) string {
	if handler, ok := r.routes[command]; !ok {
		return UnknownCommand.Error()
	} else {
		return handler(arguments)
	}
}
