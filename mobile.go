package mobile

import "github.com/gopub/types"

type IntE struct {
	Val int
	Err *Error
}

type BoolE struct {
	Val bool
	Err *Error
}

type StringE struct {
	Val string
	Err *Error
}

type Int64E struct {
	Val int64
	Err *Error
}

type ImageE struct {
	Val *types.Image
	Err *Error
}

type Float64E struct {
	Val float64
	Err *Error
}

type SecretManager interface {
	GetSecret(key string) string
	SetSecret(key, data string) bool
	DelSecret(key string) bool
}

type Uptimer interface {
	Uptime() int64
}
