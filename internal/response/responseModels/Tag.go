package responseModels

type Tag struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func NewTag() *Tag {
	return &Tag{}
}
