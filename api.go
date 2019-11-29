package main

import (
	"log"
	"fmt"
	
	"github.com/fat-max/pyro-api/route"
	"github.com/gorilla/csrf"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	log.Println("pyro api")

	r := mux.NewRouter()
	csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))

	api := r.PathPrefix("/api/v0.1").Subrouter()
	api.Use(csrfMiddleware)
	api.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "api v0.1")
	})
	api.HandleFunc("/chemicals", route.AllChemicals).Methods(http.MethodGet)
	api.HandleFunc("/chemicals/{slug}", route.Chemical).Methods(http.MethodGet)

	log.Fatalln(http.ListenAndServe(":8888", r))
}
