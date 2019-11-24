package requester

import (
	"net/http"
	"time"

	"github.com/atljoseph/api.josephgill.io/logger"
)

const (
	logID = "requester"
)

func logMessage(action string, message string) {
	logger.EntryWithMessage(logID, action, message).Info()
}

func requestFields(r *http.Request, handlerName string) logger.LogFields {
	return logger.LogFields{
		"handler": handlerName,
		"method":  r.Method,
		"uri":     r.RequestURI,
	}
}

func logRequestStart(action string, r *http.Request, handlerName string) {
	logger.Entry(logID, action).
		WithFields(requestFields(r, handlerName)).
		Info()
}

func logRequestEnd(action string, r *http.Request, handlerName string, duration time.Duration) {
	logger.Entry(logID, action).
		WithFields(requestFields(r, handlerName)).
		WithFields(logger.LogFields{
			"duration": duration,
		}).
		Info()
}
