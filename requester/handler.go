package requester

import (
	"net/http"
	"time"

	"github.com/atljoseph/api.josephgill.io/logger"
)

func requestFields(r *http.Request, handlerName string) logger.LogFields {
	return logger.LogFields{
		"handler": handlerName,
		"method":  r.Method,
		"uri":     r.RequestURI,
	}
}

// HandleWithLogging handles the request with the destination http.HandlerFunc, wrapped with logging
func HandleWithLogging(innerHandler http.Handler, name string) http.Handler {
	funcTag := "HandleWithLogging"
	funcLog := pkgLog.WithFunc(funcTag)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// can use this to log to data pipeline as well
		// or use as a centralized way to handle errors from requests

		// log request start
		funcLog.WithMessage("request start").WithFields(requestFields(r, name)).Info()

		// execute the innerHandler
		start := time.Now()
		innerHandler.ServeHTTP(w, r)
		duration := time.Since(start)

		// log stats
		funcLog.WithMessage("request end").WithFields(requestFields(r, name)).WithDuration(duration).Info()
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
