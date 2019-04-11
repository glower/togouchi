package logging

import (
	"net/http"
	"os"

	togouchi "github.com/glower/togouchi/pkg"
	"github.com/gorilla/handlers"
)

func init() {
	togouchi.Register(requestLoggingHandler)
}

func requestLoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.CombinedLoggingHandler(os.Stdout, h).ServeHTTP(w, r)
	})
}
