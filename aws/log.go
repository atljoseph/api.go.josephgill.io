package aws

import (
	"github.com/atljoseph/api.josephgill.io/logger"
)

const (
	logID = "aws"
)

func logMessage(action string, message string) {
	logger.EntryWithMessage(logID, action, message).Info()
}

func logError(action string, err error) {
	logger.EntryWithError(logID, action, err).Error()
}
