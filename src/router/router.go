package router

import (
	"portfolio-api/src/router/routes"

	"github.com/gorilla/mux"
)

// GenerateRoutes gera todas as rotas da aplicação
func GenerateRoutes() *mux.Router {
	r := mux.NewRouter()
	return routes.Config(r)
}
