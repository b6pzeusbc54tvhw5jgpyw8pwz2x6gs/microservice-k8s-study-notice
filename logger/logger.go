package logger

import (
	"github.com/onrik/logrus/filename"
	log "github.com/sirupsen/logrus"
)

func Init() {
	filenameHook := filename.NewHook()
	filenameHook.Field = "fileInfo" // Customize source field name
	log.AddHook(filenameHook)
}
