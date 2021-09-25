package exception

func NewInternalServerError(format string, a ...interface{}) APIException {
	return newException(usedNamespace, InternalServerError, format, a...)
}

func NewBadRequest(format string, a ...interface{}) APIException {
	return newException(usedNamespace, BadRequest, format, a...)
}
