package error_wrapper

type ErrorWrapper struct {
	Error     error
	ErrorCode int

	Message [2]string

	Data []interface{}

	IsMasked      bool
	MaskedMessage string
	tempMessage   string
}

func New(errorCode int, message string, isMask bool, maskedMessage string) ErrorWrapper {
	return ErrorWrapper{
		ErrorCode:     errorCode,
		IsMasked:      isMask,
		tempMessage:   message,
		MaskedMessage: maskedMessage,
	}
}
