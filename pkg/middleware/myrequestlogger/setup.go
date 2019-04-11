package myrequestlogger

import (
	"net/http"
	"os"

	togouchi "github.com/glower/togouchi/pkg"
	apachelog "github.com/lestrrat-go/apache-logformat"
)

var pattern = `User-agent is [%{User-agent}i], Request-Id is [%{X-Request-ID}i]`
var log *apachelog.ApacheLog

func init() {
	var err error
	log, err = apachelog.New(pattern)
	if err != nil {
		panic(err)
	}

	togouchi.Register(myRequestLoggingHandler)
}

func myRequestLoggingHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Wrap(h, os.Stdout).ServeHTTP(w, r)
	})
}
