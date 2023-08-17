package routes

import (
	"net/http"
	"portfolio-api/src/middlewares"

	"github.com/gorilla/mux"
)

// Route representa todas as rotas da API
type Route struct {
	Uri      string
	Method   string
	Function func(http.ResponseWriter, *http.Request)
	Auth     bool
}

func Config(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, loginRoute)
	routes = append(routes, postsRoutes...)

	for _, route := range routes {

		if route.Auth {
			r.HandleFunc(route.Uri, middlewares.Logger(middlewares.Auth(route.Function))).Methods(route.Method)
		} else {
			r.HandleFunc(route.Uri, middlewares.Logger(route.Function)).Methods(route.Method)
		}
	}

	return r
}
