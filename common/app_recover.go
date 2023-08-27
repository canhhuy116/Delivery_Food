package common

import (
	"log"
	"runtime/debug"
)

func AppRecover() {
	if r := recover(); r != nil {
		log.Println("Recovered from panic:", r)
		debug.PrintStack()
		// You can add more handling logic here, such as logging the panic or notifying administrators.
	}
}
