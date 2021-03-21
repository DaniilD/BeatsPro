package Tag

type Tag struct {
	Id        int
	Title     string
	IsDeleted bool
}

func NewTag() *Tag {
	return &Tag{}
}
