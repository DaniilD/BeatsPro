package repositories

import (
	"BeatsPro/internal/db"
	"BeatsPro/internal/models/Track"
	"fmt"
)

type TrackRepository struct {
	store db.Store
}

func NewTrackRepository(store db.Store) *TrackRepository {
	return &TrackRepository{
		store: store,
	}
}

func (trackRepository *TrackRepository) CreateTrack(track *Track.Track) (int, error) {
	var id int64
	sql := "INSERT INTO %s " +
		"(`userId` `name`, `description`, `duration`, `bpm`, `isDeleted`) " +
		"VALUES " +
		"(?, '?', '?', ?, ?, ?);"

	query := fmt.Sprintf(sql, TRACKS_TABLE)
	result, err := trackRepository.store.Exec(query,
		track.UserId,
		track.Name,
		track.Description,
		track.Duration,
		track.Bpm,
		track.IsDeleted)

	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (trackRepository *TrackRepository) UpdateTrack(track *Track.Track) error {
	sql := "UPDATE %s " +
		"SET " +
		"`userId` = ?, " +
		"`name` = '?', " +
		"'description' = ?, " +
		"`duration` = ? " +
		"`bpm` = ? " +
		"`isDeleted` = ?" +
		"WHERE `id` = ?"

	query := fmt.Sprintf(sql, TRACKS_TABLE)

	_, err := trackRepository.store.Exec(query,
		track.UserId,
		track.Name,
		track.Description,
		track.Duration,
		track.Bpm,
		track.IsDeleted,
		track.Id)

	if err != nil {
		return err
	}

	return nil
}

func (trackRepository *TrackRepository) GetById(id int) (*Track.Track, error) {
	sql := "SELECT " +
		"`id`, `userId`, `name`, `description` " +
		"`dateTimeCreation`, `duration`, `bpm`, `isDeleted` " +
		"FROM %s" +
		"WHERE `id` = ?"

	query := fmt.Sprintf(sql, TRACKS_TABLE)
	row := trackRepository.store.QueryRow(query, id)
	track := Track.NewTrack()
	err := row.Scan(track.Id,
		track.UserId,
		track.Name,
		track.Description,
		track.DateTimeCreation,
		track.Duration,
		track.Bpm,
		track.IsDeleted)

	if err != nil {
		return nil, err
	}

	return track, nil
}
