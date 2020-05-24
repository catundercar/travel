package main

import (
	"log"
	"net/http"
	"time"
)

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func timeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()

		next.ServeHTTP(w, r)

		timeElapsed := time.Since(timeStart)
		log.Println(timeElapsed)
	})
}

func main() {
	//mux := http.NewServeMux()
	//	mux.Handle("/", timeMiddleware(http.HandlerFunc(hello)))

	r := NewRouter()
	//r.h = http.HandlerFunc(hello)
	r.Use(timeMiddleware)
	r.Add("/hello", http.HandlerFunc(hello))
	http.Handle("/", r)

	http.ListenAndServe(":8999", nil)
}
