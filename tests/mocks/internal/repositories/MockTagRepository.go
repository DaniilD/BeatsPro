package repositories_mock

import (
	"BeatsPro/internal/db"
	"BeatsPro/internal/models/Tag"
	"github.com/stretchr/testify/mock"
)

type MockTagRepository struct {
	mock.Mock
	store db.Store
}

func (mockTagRepository *MockTagRepository) CreateTag(tag *Tag.Tag) (int, error) {
	args := mockTagRepository.Called(tag)
	return args.Int(tag.Id), args.Error(0)
}

func (mockTagRepository *MockTagRepository) UpdateTag(tag *Tag.Tag) error {
	args := mockTagRepository.Called(tag)

	return args.Error(0)
}

func (mockTagRepository *MockTagRepository) GetById(id int) (*Tag.Tag, error) {
	args := mockTagRepository.Called(id)

	tag := Tag.NewTag()
	tag.Id = id

	return tag, args.Error(0)
}
