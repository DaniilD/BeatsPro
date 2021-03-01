package db

import (
	"BeatsPro/internal/configs"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
)

type MySqlStore struct {
	config *configs.DBConfig
	db     *sqlx.DB
}

func NewMysqlStore(config *configs.DBConfig) *MySqlStore {
	mySqlStore := &MySqlStore{
		config: config,
	}
	if err := mySqlStore.open(); err != nil {
		log.Fatal(err)
	}
	return mySqlStore
}

func (store *MySqlStore) open() error {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		store.config.User, store.config.Password, store.config.Host, fmt.Sprintf("%v", store.config.Port), store.config.Name)

	db, err := sqlx.Open(store.config.Driver, connectionString)

	if err != nil {
		return err
	}
	store.db = db

	return nil
}

func (store *MySqlStore) QueryRow(query string, args ...interface{}) *sql.Row {
	return store.db.QueryRow(query, args...)
}
