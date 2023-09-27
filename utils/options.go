package utils

type Handler interface {
	Name() string
	Handle()
}

var handlers = make(map[string]Handler)

func AddHandler(handler Handler) {
	handlers[handler.Name()] = handler
}

func RunHandlers() {
	for _, v := range handlers {
		v.Handle()
	}
}
