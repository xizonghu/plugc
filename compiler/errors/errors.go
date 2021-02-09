package errors

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}


func New(text string) error {
	panic(text)
	return &errorString{text}
}

