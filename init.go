package netconn

import "log"

var loggerObj logger

func init() {
	// use default log
	loggerObj = loggerFunc(func(msg string) {
		log.Print(msg)
	})
}

// SetLogger ...
func SetLogger(fn func(msg string)) {
	// use custom log
	loggerObj = loggerFunc(func(msg string) {
		fn(msg)
	})
}
