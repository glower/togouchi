package main

import (
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"
)

func timeoutHandler(h http.Handler) http.Handler {
	return http.TimeoutHandler(h, 1*time.Second, "timed out")
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Logged connection from %s", r.RemoteAddr)
		next.ServeHTTP(w, r)
	})
}

func withTracing(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Tracing request for %s", r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

func myApp(w http.ResponseWriter, r *http.Request) {
	log.Println("Hello world!")
	w.Write([]byte("Hello world!\n"))
}

func main() {
	m := http.NewServeMux()
	m.HandleFunc("/", myApp)
	chain := alice.New(timeoutHandler, withLogging, withTracing).Then(m)
	http.ListenAndServe(":8000", chain)
}
