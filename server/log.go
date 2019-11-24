package server

import (
	"github.com/atljoseph/api.josephgill.io/logger"
)

const (
	logID = "server"
)

func logMessage(action string, message string) {
	logger.EntryWithMessage(logID, action, message).Info()
}
