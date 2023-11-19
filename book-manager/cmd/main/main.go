package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"pro1/book-manager/pkg/routes"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}
