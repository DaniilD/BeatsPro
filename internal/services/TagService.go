package services

import (
	"BeatsPro/internal/models/Tag"
	"BeatsPro/internal/repositories"
)

type TagService struct {
	tagRepository *repositories.TagRepository
}

func NewTagService(repository *repositories.TagRepository) *TagService {
	return &TagService{
		tagRepository: repository,
	}
}

func (tagService *TagService) CreateTag(tag *Tag.Tag) (int, error) {
	return tagService.tagRepository.CreateTag(tag)
}

func (tagService *TagService) UpdateTag(tag *Tag.Tag) error {
	return tagService.tagRepository.UpdateTag(tag)
}

func (tagService *TagService) GetById(id int) (*Tag.Tag, error) {
	return tagService.tagRepository.GetById(id)
}
