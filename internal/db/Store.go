package db

type Store interface {
	open() error
}
