package storage

type Storage interface {
	Connect() error
	Ping() error
	Close() error
}
