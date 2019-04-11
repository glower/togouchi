package middleware

import (
	// plug in the standard middleware
	_ "github.com/glower/togouchi/pkg/middleware/logging"
	_ "github.com/glower/togouchi/pkg/middleware/myrequestlogger"
	_ "github.com/glower/togouchi/pkg/middleware/requestid"
)
