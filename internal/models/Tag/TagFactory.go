package Tag

type TagFactory struct {
}

func NewTagFactory() *TagFactory {
	return &TagFactory{}
}

func (tagFactory *TagFactory) Make(title string) *Tag {
	tag := NewTag()
	tag.Title = title

	return tag
}
