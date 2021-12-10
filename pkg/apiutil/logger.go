package apiutil

// Logger is an abstracted logger type.
type Logger interface {
	Println(string, ...interface{})
}
