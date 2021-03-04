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
	var id int64
	query := fmt.Sprintf("INSERT INTO %s (title) VALUES (?);", TAGS_TABLE)
	result, err := tagRepository.store.Exec(query, tag.Title)

	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (tagRepository *TagRepository) UpdateTag(tag *Tag.Tag) error {
	query := fmt.Sprintf("UPDATE %s SET `title` = ? WHERE `id` = ?;", TAGS_TABLE)
	_, err := tagRepository.store.Exec(query, tag.Title, tag.Id)

	if err != nil {
		return err
	}

	return nil
}

func (tagRepository *TagRepository) GetById(id int) (*Tag.Tag, error) {
	query := fmt.Sprintf("SELECT `id`, `title` FROM %s WHERE `id` = ?;", TAGS_TABLE)
	row := tagRepository.store.QueryRow(query, id)
	tag := Tag.NewTag()
	err := row.Scan(tag.Id, tag.Title)

	if err != nil {
		return nil, err
	}

	return tag, nil
}
