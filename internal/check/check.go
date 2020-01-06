package check

import (
	"log"
)

//Error export
func Error(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
