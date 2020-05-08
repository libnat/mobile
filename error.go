package mobile

import "fmt"

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) String() string {
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}

func ToError(err error) *Error {
	if err == nil {
		return nil
	}

	for {
		u, ok := err.(interface{ Unwrap() error })
		if !ok {
			break
		}
		err = u.Unwrap()
	}

	if ge, ok := err.(*Error); ok {
		return NewError(ge.Code, ge.Message)
	}
	return NewError(0, err.Error())
}
