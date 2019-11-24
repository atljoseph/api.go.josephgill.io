package handlers

import (
	"github.com/atljoseph/api.josephgill.io/logger"
)

const (
	logID = "handlers"
)

func logMessage(action string, message string) {
	logger.EntryWithMessage(logID, action, message).Info()
}
