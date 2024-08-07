package models

type Response struct {
	HttpCode int
	Body     []byte
	Headers  map[string]string
}
