package Tag

type Tag struct {
	id    int
	title string
}

func NewTag() *Tag {
	return &Tag{}
}
