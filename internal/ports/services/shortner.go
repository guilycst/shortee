package services

type Shortener interface {
	Shorten(url string) (string, error)
	Resolve(id string) (string, error)
}
