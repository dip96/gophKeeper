package config

type ErrorConfig struct {
	msg   string
	error error
}

func New(text string, err error) error {
	return &ErrorConfig{text, err}
}

func (e *ErrorConfig) Error() string {
	return e.msg
}
