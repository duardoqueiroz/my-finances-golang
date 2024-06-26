package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type Database struct {
	user     string
	password string
	port     string
	name     string
	host     string
	conn     *sqlx.DB
}

func SetupCredentials(user, password, port, name, host string) *Database {
	return &Database{
		user:     user,
		password: password,
		port:     port,
		name:     name,
		host:     host,
	}
}

func getConnString(user, password, port, name, host string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, name, host, port)
}

func (c *Database) Connect() error {
	connectionString := getConnString(c.user, c.password, c.port, c.name, c.host)
	conn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *Database) Disconnect() error {
	return c.conn.Close()
}

func (c *Database) Connection() *sqlx.DB {
	return c.conn
}
