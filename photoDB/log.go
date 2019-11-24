package photoDB

import "github.com/atljoseph/api.josephgill.io/logger"

const (
	logID = "photoDB"
)

func logMessage(action string, message string) {
	logger.EntryWithMessage(logID, action, message).Info()
}
