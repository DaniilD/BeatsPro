package consumerError

type NotFoundErrorResponse struct {
	Status string `json:"status"`
	Error  *Error `json:"error"`
}

func NewNotFoundErrorResponse(status string, error *Error) *NotFoundErrorResponse {
	return &NotFoundErrorResponse{
		Status: status,
		Error:  error,
	}
}
