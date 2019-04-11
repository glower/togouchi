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
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/products/{id}/", ProductsHandler)
	r.HandleFunc("/articles/{id}/", ArticlesHandler)
	http.Handle("/", r)

	log.Fatal(http.ListenAndServe(":8000", togouchi.Run(r)))
}

// HomeHandler ...
func HomeHandler(res http.ResponseWriter, req *http.Request) {
	rid := req.Header.Get("X-Request-ID")
	log.Printf(">>> HomeHandler(): RID: %s\n", rid)
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
