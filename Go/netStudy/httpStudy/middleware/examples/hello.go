package main

import (
	"log"
	"net/http"
	"time"

	"go.study.org/netStudy/httpStudy/middleware"
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
	mux := http.NewServeMux()

	hm := middleware.NewHTTPMiddleware()
	hm.Use(timeMiddleware)
	hm.Add("/hello", http.HandlerFunc(hello))
	hm.Register(mux)

	http.ListenAndServe(":8999", mux)
}
