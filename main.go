package main

import (
	"fmt"
	"log"
	"net/http"
	"portfolio-api/src/config"
	"portfolio-api/src/router"
)

func main() {
	config.Initialize()

	r := router.GenerateRoutes()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
