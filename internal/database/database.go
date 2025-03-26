package database

type Database interface {
	Connect() error
	Disconnect() error
}
