package apierr

import (
	"github.com/atljoseph/api.josephgill.io/logger"
)

const (
	logID = "apierr"
)

// TODO: remove error logging from here?

func logError(err StackedError) {
	logger.EntryWithError(logID, "error", err).Warn()
}
