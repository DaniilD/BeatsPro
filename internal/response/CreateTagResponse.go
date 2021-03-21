package response

type CreateTagResponse struct {
	Status string `json:"status"`
	Data   struct {
		Id int `json:"id"`
	} `json:"data"`
}

func NewCreateTagResponse() *CreateTagResponse {
	return &CreateTagResponse{}
}
