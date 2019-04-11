package togouchi

import (
	"net/http"
	"sort"
	"sync"
)

var (
	middlewaresM sync.RWMutex
	middlewares  []Middleware
)

// HandlerCall is a http handler for the middleware
type HandlerCall func(http.Handler) http.Handler

// Middleware ...
type Middleware struct {
	HandlerCall
	Description string
	Order       int
}

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

	sort.Slice(middlewares, func(i, j int) bool {
		return middlewares[i].Order > middlewares[j].Order
	})

	for _, m := range middlewares {
		handler := m.HandlerCall
		h = handler(h)
	}

	return h
}

// chainMiddleware provides syntactic sugar to create a new middleware
// which will be the result of chaining the ones received as parameters.
// func chainMiddleware(middlewares []Middleware) HandlerCall {
// 	return func(final http.HandlerFunc) http.HandlerFunc {
// 		return func(w http.ResponseWriter, r *http.Request) {
// 			last := final
// 			for _, m := range middlewares {
// 				handler := m.HandlerCall
// 				last = handler(last)
// 			}
// 			last(w, r)
// 		}
// 	}
// }
