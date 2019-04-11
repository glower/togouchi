package requestid

import (
	"net/http"

	"github.com/ascarter/requestid"
	togouchi "github.com/glower/togouchi/pkg"
)

func init() {
	togouchi.Register(requestIDTracing)
}

func requestIDTracing(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requestid.RequestIDHandler(h).ServeHTTP(w, r)
	})
}
