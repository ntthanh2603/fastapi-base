package response

import "fmt"

type APIError struct {
	StatusCode int
	Message    string
	Err        interface{}
}

func (e *APIError) Error() string {
	switch v := e.Err.(type) {
	case string:
		return v
	case error:
		return v.Error()
	default:
		return fmt.Sprintf("%v", v)
	}
}

func NewAPIError(status int, message string, err interface{}) *APIError {
	return &APIError{
		StatusCode: status,
		Message:    message,
		Err:        err,
	}
}
