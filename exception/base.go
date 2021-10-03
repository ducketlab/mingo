package exception

import "fmt"

type APIException interface {
	error
	ErrorCode() int
	WithMeta(m interface{}) APIException
	Meta() interface{}
	WithData(d interface{}) APIException
	Data() interface{}
	Is(code int) bool
	Namespace() string
	Reason() string
}

func newException(namespace Namespace, code int, format string, a ...interface{}) *exception {
	return &exception{
		namespace: namespace,
		code:      code,
		reason:    codeReason(code),
		message:   fmt.Sprintf(format, a...),
	}
}

type exception struct {
	namespace Namespace
	code      int
	reason    string
	message   string
	meta      interface{}
	data      interface{}
}

func (e *exception) Error() string {
	return e.message
}

func (e *exception) ErrorCode() int {
	return e.code
}

func (e *exception) WithMeta(m interface{}) APIException {
	e.meta = m
	return e
}

func (e *exception) Meta() interface{} {
	return e.meta
}

func (e *exception) WithData(d interface{}) APIException {
	e.data = d
	return e
}

func (e *exception) Data() interface{} {
	return e.data
}

func (e *exception) Is(code int) bool {
	return e.ErrorCode() == code
}

func (e *exception) Namespace() string {
	return e.namespace.String()
}

func (e *exception) Reason() string {
	return e.reason
}
