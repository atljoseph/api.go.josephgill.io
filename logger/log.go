package logger

const (
	logID = "logger"
)

func logMessage(action string, message string) {
	EntryWithMessage(logID, action, message).Info()
}
