package postgres

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type database struct {
	user     string
	password string
	port     string
	name     string
	host     string
	conn     *sqlx.DB
}

func NewPostgresDB() (*database, error) {
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	port := os.Getenv("POSTGRES_PORT")
	name := os.Getenv("POSTGRES_DB")
	host := os.Getenv("POSTGRES_HOST")

	err := validateCredentials(user, password, port, name, host)
	if err != nil {
		return nil, err
	}

	return &database{
		user:     user,
		password: password,
		port:     port,
		name:     name,
		host:     host,
	}, nil
}

func validateCredentials(user, password, port, name, host string) error {
	if user == "" {
		return fmt.Errorf("user is required")
	}
	if password == "" {
		return fmt.Errorf("password is required")
	}
	if port == "" {
		return fmt.Errorf("port is required")
	}
	if name == "" {
		return fmt.Errorf("name is required")
	}
	if host == "" {
		return fmt.Errorf("host is required")
	}
	return nil
}

func getConnString(user, password, port, name, host string) string {
	return fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", user, password, name, host, port)
}

func (c *database) Connect() error {
	connectionString := getConnString(c.user, c.password, c.port, c.name, c.host)
	conn, err := sqlx.Connect("postgres", connectionString)
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *database) Disconnect() error {
	return c.conn.Close()
}

func (c *database) Connection() *sqlx.DB {
	return c.conn
}
