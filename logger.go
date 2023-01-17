package netconn

import "fmt"

type logger interface {
	print(string)
	printf(string, ...any)
}

type loggerFunc func(string)

func (lF loggerFunc) print(msg string) {
	lF(msg)
}
func (lF loggerFunc) printf(msg string, args ...any) {
	lF(fmt.Sprintf(msg, args...))
}
