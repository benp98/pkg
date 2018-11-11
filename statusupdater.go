package pkg

import (
	"log"
)

// StatusUpdater receives messages of the package management progress
type StatusUpdater interface {
	Message(string)
}

// LogStatusUpdater prints the package management status via the go log package
type LogStatusUpdater struct{}

// Message writes the message to the default log output
func (statusUpater LogStatusUpdater) Message(message string) {
	log.Println(message)
}
