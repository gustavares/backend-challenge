package main

import (
	"flag"
	"net/http"
	"tasks/api"
	"tasks/config"
	"tasks/datastore"
)

func main() {
	c := config.New()

	// db connection
	db := datastore.New(c)

	// setting up routes
	h := api.NewHandlers(db)
	r := h.RegisterRoutes()

	port := flag.String("port", ":3001", "to-do")
	http.ListenAndServe(*port, r)
}
