package exception

import "fmt"

func NewApiException(namespace string, code int, reason, format string, a ...interface{}) APIException {

	if code == 0 {
		code = -1
	}

	return &exception{
		namespace: Namespace(namespace),
		code:      code,
		reason:    codeReason(code),
		message:   fmt.Sprintf(format, a...),
	}
}

func NewInternalServerError(format string, a ...interface{}) APIException {
	return newException(usedNamespace, InternalServerError, format, a...)
}

func NewBadRequest(format string, a ...interface{}) APIException {
	return newException(usedNamespace, BadRequest, format, a...)
}
