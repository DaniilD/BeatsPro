package repositories

import (
	"BeatsPro/internal/db"
	"BeatsPro/internal/models/User"
	"fmt"
)

type UserRepository struct {
	store db.Store
}

func NewUserRepository(store db.Store) *UserRepository {
	return &UserRepository{
		store: store,
	}
}

func (userRepository *UserRepository) CreateUser(user *User.User) (int, error) {
	var id int64
	sql := "INSERT INTO %s " +
		"(`login`, `password`, `email`, `name`, `lastName`, `type`, `dateOfBirth`, " +
		"`isDeleted`, `isBanned`, `isConfirmed`) " +
		"VALUES " +
		"('?', '?', '?', '?', '?', ?, '?', ?, ?, ?);"

	query := fmt.Sprint(sql, USERS_TABLE)
	result, err := userRepository.store.Exec(query)

	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()

	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (userRepository *UserRepository) UpdateUser(user *User.User) error {
	sql := "UPDATE %s SET " +
		"`login` = '?', " +
		"`password` = '?', " +
		"`email` = '?', " +
		"`name` = '?', " +
		"`lastName` = '?', " +
		"`type` = ?, " +
		"`dateTimeCreate` = '?', " +
		"`dateOfBirth` = '?', " +
		"`isDeleted` = ?, " +
		"`isBanned` = ?, " +
		"`isConfirmed` = ? " +
		"WHERE `id` = ?;"

	query := fmt.Sprintf(sql, USERS_TABLE)

	_, err := userRepository.store.Exec(query,
		user.Login,
		user.Password,
		user.Email,
		user.Name,
		user.LastName,
		user.Type,
		user.DateTimeCreation,
		user.DateOfBirth,
		user.IsDeleted,
		user.IsBanned,
		user.IsConfirmed)

	if err != nil {
		return err
	}

	return nil
}

// dateTimeCreate -> DateTimeCreation
func (userRepository *UserRepository) GetById(id int) (*User.User, error) {
	sql := "SELECT " +
		"`id`, `login`, `password`, `email`, `name`, `lastName`, `type`, " +
		"`dateTimeCreate`, `dateOfBirth`, `isDeleted`, `isBanned`, `isConfirmed` " +
		"FROM %s " +
		"WHERE id = ?;"

	query := fmt.Sprintf(sql, USERS_TABLE)
	row := userRepository.store.QueryRow(query, id)
	user := User.NewUser()
	err := row.Scan(
		user.Id,
		user.Login,
		user.Password,
		user.Email,
		user.Name,
		user.LastName,
		user.Type,
		user.DateTimeCreation,
		user.DateOfBirth,
		user.IsDeleted,
		user.IsBanned,
		user.IsConfirmed)

	if err != nil {
		return nil, err
	}

	return user, nil
}
