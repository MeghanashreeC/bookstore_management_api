package main

import (
	"log"
	"net/http"

	"github.com/MeghanashreeC/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

// to run the server
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
