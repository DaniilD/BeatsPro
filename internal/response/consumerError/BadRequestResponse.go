package consumerError

type BadRequestResponse struct {
	Status string `json:"status"`
	Error  *Error `json:"error"`
}

func NewBadRequestResponse(status string, error *Error) *BadRequestResponse {
	return &BadRequestResponse{
		Status: status,
		Error:  error,
	}
}
