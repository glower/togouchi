package main

import (
	"fmt"
	"log"
	"net/http"

	// load all registerd middleware

	togouchi "github.com/glower/togouchi/pkg"
	_ "github.com/glower/togouchi/pkg/middleware"
	"github.com/gorilla/mux"
)

func main() {
	m := http.NewServeMux()

	m.HandleFunc("/", HomeHandler)
	m.HandleFunc("/products/{id}/", ProductsHandler)
	m.HandleFunc("/articles/{id}/", ArticlesHandler)

	log.Fatal(http.ListenAndServe(":8000", togouchi.Run(m)))
}

// HomeHandler ...
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("HomeHandler(): Welcome Home!")
	res.WriteHeader(http.StatusOK)
	fmt.Fprint(res, "Welcome Home!\n")
}

// ArticlesHandler ...
func ArticlesHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Article: %v\n", vars["id"])
}

// ProductsHandler ...
func ProductsHandler(res http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	res.WriteHeader(http.StatusOK)
	fmt.Fprintf(res, "Product: %v\n", vars["id"])
}
