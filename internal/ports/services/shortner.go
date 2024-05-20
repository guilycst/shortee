package services

type ErrEmptyURL struct {
}

func (e *ErrEmptyURL) Error() string {
	return "empty url"
}

type ErrURLTooLong struct {
}

func (e *ErrURLTooLong) Error() string {
	return "url too long"
}

type Shortener interface {
	Shorten(url string) (string, error)
	Resolve(id string) (string, error)
}
