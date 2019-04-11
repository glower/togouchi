package myrequestlogger

import (
	"net/http"
	"os"

	togouchi "github.com/glower/togouchi/pkg"
	apachelog "github.com/lestrrat-go/apache-logformat"
)

var pattern = `User-agent is [%{User-agent}i], Request-Id is [%{X-Request-ID}i]`
var alog *apachelog.ApacheLog

const order = 200

func init() {
	var err error
	alog, err = apachelog.New(pattern)
	if err != nil {
		panic(err)
	}

	togouchi.Register(togouchi.Middleware{
		HandlerCall: myRequestLoggingHandler,
		Description: "my request logger",
		Order:       order,
	})
}

func myRequestLoggingHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		alog.Wrap(next, os.Stdout).ServeHTTP(w, r)
	})
}
