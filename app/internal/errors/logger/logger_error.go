package logger

type ErrorLogger struct {
	msg   string
	error error
}

func New(text string, err error) error {
	return &ErrorLogger{text, err}
}

func (e *ErrorLogger) Error() string {
	return e.msg
}
