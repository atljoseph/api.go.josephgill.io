package handlers

import (
	"github.com/atljoseph/api.josephgill.io/logger"
)

var pkgLog *logger.Log

func init() {
	pkgLog = logger.ForPackage("handlers")
}
