package Tag

type Tag struct {
	Id    int
	Title string
}

func NewTag() *Tag {
	return &Tag{}
}
