package requestid

import (
	"github.com/ascarter/requestid"
	togouchi "github.com/glower/togouchi/pkg"
)

const order = 1

func init() {
	togouchi.Register(togouchi.Middleware{
		HandlerCall: requestid.RequestIDHandler,
		Description: "Request-ID tracker",
		Order:       order,
	})
}
