package middleware

import (
	"net/http"
)

type middleware func(http.Handler) http.Handler

// A HTTPMux registers the handler for the given pattern.
type HTTPMux interface {
	Handle(pattern string, handler http.Handler)
}

// A HTTPMiddleware represents a middleware of http.
type HTTPMiddleware struct {
	// middlewareChain is a chain of http-middleware.
	middlewareChain []middleware
	mux             map[string]http.Handler
}

// NewHTTPMiddleware returns a HTTPMiddleware.
func NewHTTPMiddleware() *HTTPMiddleware {
	return &HTTPMiddleware{
		mux: make(map[string]http.Handler),
	}
}

// Use will append a middleware to the middlewareChain.
func (hm *HTTPMiddleware) Use(m middleware) {
	hm.middlewareChain = append(hm.middlewareChain, m)
}

// Add will register the middlewareChain for a Handler.
func (hm *HTTPMiddleware) Add(pattern string, h http.Handler) {
	var mergedHandler = h

	for i := len(hm.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = hm.middlewareChain[i](mergedHandler)
	}

	hm.mux[pattern] = mergedHandler
}

// Register registers mux to a http mux.
func (hm *HTTPMiddleware) Register(mux HTTPMux) {
	for p, h := range hm.mux {
		mux.Handle(p, h)
	}
}
