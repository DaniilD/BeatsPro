package operationErrors

type ValidateError struct {
	message string
}

func NewValidateError(message string) *ValidateError {
	return &ValidateError{
		message: message,
	}
}

func (error *ValidateError) Error() string {
	return error.message
}
