package routes

import (
	"net/http"
	"portfolio-api/src/router/controllers"
)

var loginRoute = Route{
	Uri:      "/login",
	Method:   http.MethodPost,
	Function: controllers.Login,
	Auth:     false,
}
