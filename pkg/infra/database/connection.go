package database

type Connection interface {
	Connect() error
	Disconnect() error
}
