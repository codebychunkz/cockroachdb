package main

import "log"

type Closable func() error

func LoggingClose(close Closable) {
	err := close()
	if err != nil {
		log.Printf("Error while closing: %s", err)
	}
}
