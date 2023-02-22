package inits

import (
	log "WindowsListener/models"
	nested "github.com/antonfisher/nested-logrus-formatter"
)

func LogInit() {
	log.Logger.SetFormatter(&nested.Formatter{
		TimestampFormat: "2006-01-01 10:06:32",
		HideKeys:        true,
		TrimMessages:    true,
		FieldsOrder:     nil,
	})
}
