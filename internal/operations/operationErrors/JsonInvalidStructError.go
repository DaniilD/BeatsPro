package operationErrors

type JsonInvalidStructError struct {
}

func (error *JsonInvalidStructError) Error() string {
	return "Invalid JSON structure"
}
