package Tag

type TagFactory struct {
}

func NewTagFactory() *TagFactory {
	return &TagFactory{}
}

func (tagFactory *TagFactory) Make(title string, isDeleted bool) *Tag {
	tag := NewTag()
	tag.Title = title
	tag.IsDeleted = isDeleted

	return tag
}
