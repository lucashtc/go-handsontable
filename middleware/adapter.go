package middleware

import "net/http"

// Adapter ...
type Adapter func(h http.Handler) http.Handler

// Adapt middleware
func Adapt(h http.Handler, adapters ...Adapter) http.Handler {
	for _, adapter := range adapters {
		h = adapter(h)
	}
	return h
}