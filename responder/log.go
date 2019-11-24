package responder

import (
	"github.com/atljoseph/api.josephgill.io/logger"
)

const (
	logID = "responder"
)

func logErrorResponse(action string, errText string, statusCode int) {
	logger.Entry(logID, action).
		WithFields(logger.LogFields{
			"error": errText,
			"code":  statusCode,
		}).Error()
}
