package main

import (
	"net/http"

	"github.com/andrii-stasiuk/go-testing/bdd-testing/lib"
)

func main() {
	routers := lib.SetUserRoutes()
	http.ListenAndServe(":8080", routers)
}
