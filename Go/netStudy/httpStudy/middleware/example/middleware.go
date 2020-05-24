package main

import (
	"net/http"
)

type middleware func(http.Handler) http.Handler

type muxEntry struct {
	pattern string
	handler http.Handler
}

type Router struct {
	middlewareChain []middleware
	mux             map[string]muxEntry
	//h               http.Handler
}

func NewRouter() *Router {
	return &Router{
		mux: make(map[string]http.Handler),
		//h: http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {}),
	}
}

func (r *Router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
	//r.h = m(r.h)
}

func (r *Router) Add(pattern string, h http.Handler) {
	var mergedHandler = h

	for i := len(r.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = r.middlewareChain[i](mergedHandler)
	}

	r.mux[pattern] = muxEntry{pattern: route, handler: mergedHandler}
}

//func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
//	//r.h.ServeHTTP(w, req)
//	//h.ServeHTTP(w, req)
//	for prefix, h := range r.mux {
//		if strings.HasPrefix(req.URL.Path, prefix) {
//			h.ServeHTTP(w, req)
//		}
//	}
//}
