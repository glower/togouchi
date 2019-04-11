package togouchi

import (
	"net/http"
	"sync"
)

var (
	middlewaresM sync.RWMutex
	middlewares  []Middleware
)

// Middleware is a http handler for the middleware
type Middleware func(http.Handler) http.Handler

// Register the middleware plugin
func Register(m Middleware) {
	middlewaresM.Lock()
	defer middlewaresM.Unlock()

	middlewares = append(middlewares, m)
}

// Run ...
func Run(h http.Handler) http.Handler {
	if h == nil {
		h = http.DefaultServeMux
	}

	for i := range middlewares {
		h = middlewares[len(middlewares)-1-i](h)
	}

	return h
}
