package requester

import (
	"net/http"
	"time"
)

// HandleWithLogging handles the request with the destination http.HandlerFunc, wrapped with logging
func HandleWithLogging(innerHandler http.Handler, name string) http.Handler {
	funcTag := "HandleWithLogging"
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// can use this to log to data pipeline as well
		// or use as a centralized way to handle errors from requests

		// log request start
		logRequestStart(funcTag, r, name)

		// execute the innerHandler
		start := time.Now()
		innerHandler.ServeHTTP(w, r)
		duration := time.Since(start)

		// log stats
		logRequestEnd(funcTag, r, name, duration)
	})
}

// // HandlerWithContext wraps a http.HandlerFunc by adding extra context to the request
// func HandlerWithContext(h http.Handler, pathParams []string) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		// Take the context out from the request
// 		ctx := r.Context()

// 		// Get new context with the desired key-value added
// 		ctx = context.WithValue(ctx, contextPathParamsKey, pathParams)

// 		// Get new http.Request with the new context
// 		r = r.WithContext(ctx)

// 		// Call your original http.Handler
// 		h.ServeHTTP(w, r)
// 	})
// }
