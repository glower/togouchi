package logging

import (
	"net/http"
	"os"

	togouchi "github.com/glower/togouchi/pkg"
	"github.com/gorilla/handlers"
)

const order = 100

func init() {
	togouchi.Register(togouchi.Middleware{
		HandlerCall: requestLoggingHandler,
		Description: "Apache Combined Logging Handler from gorilla",
		Order:       order,
	})
}

func requestLoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.CombinedLoggingHandler(os.Stdout, next).ServeHTTP(w, r)
	})
}
