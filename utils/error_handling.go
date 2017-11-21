package utils

import (
	log "github.com/sirupsen/logrus"
)


func CheckFatal(err error) {
	if(err != nil) {
		log.WithField("error", err).Fatal("Fatal error.")
	}
}
