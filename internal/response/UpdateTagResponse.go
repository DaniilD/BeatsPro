package response

type UpdateTagResponse struct {
	Status string `json:"status"`
}

func NewUpdateTagResponse() *UpdateTagResponse {
	return &UpdateTagResponse{}
}
