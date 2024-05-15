package services

type Closeable interface {
	Close() error
}
