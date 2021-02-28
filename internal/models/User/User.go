package User

import "time"

type User struct {
	Id               int
	Login            string
	Password         string
	Email            string
	Name             string
	LastName         string
	Type             int
	DateTimeCreation time.Time
	DateOfBirth      time.Time
	IdDeleted        string
	IsBanned         string
	IsConfirmed      string
}

func NewUser() *User {
	return &User{}
}
