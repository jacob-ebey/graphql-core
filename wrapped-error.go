package core

type WrappedError struct {
	Message       string
	Code          string
	InternalError error
}

const DefaultErrorMessage = "Unknown error."

func (err *WrappedError) Error() string {
	if err.Message != "" {
		return err.Message
	}

	if err.InternalError != nil {
		return err.InternalError.Error()
	}

	return DefaultErrorMessage
}

func (e *WrappedError) Extensions() map[string]interface{} {
	var internalError *string
	if e.InternalError != nil {
		message := e.InternalError.Error()
		internalError = &message
	}

	var code *string
	if e.Code != "" {
		code = &e.Code
	}

	return map[string]interface{}{
		"internalError": internalError,
		"code":          code,
	}
}
