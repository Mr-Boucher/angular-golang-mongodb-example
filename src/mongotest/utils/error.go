package utils

type Error interface {
	Error() string
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

// New returns an error that formats as the given text.
func NewError(text string) error {
	return &errorString{text}
}

func (e *errorString) Error() string {
	return e.s
}