package response

type SuccessResponse struct {
	Status string `json:"status"`
}

func NewSuccessResponse() *SuccessResponse {
	return &SuccessResponse{
		Status: STATUS_SUCCESS,
	}
}
