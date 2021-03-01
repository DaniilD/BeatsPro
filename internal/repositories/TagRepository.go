package repositories

import (
	"BeatsPro/internal/db"
	"BeatsPro/internal/models/Tag"
	"fmt"
)

type TagRepository struct {
	store db.Store
}

func NewTagRepository(store db.Store) *TagRepository {
	return &TagRepository{
		store: store,
	}
}

func (tagRepository *TagRepository) CreateTag(tag *Tag.Tag) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (title) VALUES (?) RETURNING `id`", TAGS_TABLE)
	row := tagRepository.store.QueryRow(query, tag.Title)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (tagRepository *TagRepository) UpdateTag(tag *Tag.Tag) error {
	return nil
}

func (tagRepository *TagRepository) GetById(id int) *Tag.Tag {
	return nil
}
