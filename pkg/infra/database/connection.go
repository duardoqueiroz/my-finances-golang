package database

import "github.com/jmoiron/sqlx"

type Connection interface {
	Connect() error
	Disconnect() error
	Connection() *sqlx.DB
}
