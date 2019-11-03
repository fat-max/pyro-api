package main

import (
	_ "bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/fat-max/pyro-api/route"
	"github.com/gorilla/mux"
)

func main() {
	log.Println("pyro api")

	r := mux.NewRouter()
	api := r.PathPrefix("/api/v0.1").Subrouter()
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v0.1")
	})

	api.HandleFunc("/chemicals", route.AllChemicals).Methods(http.MethodGet)
	api.HandleFunc("/chemicals/{id}", route.Chemical).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe(":8080", r))
}
