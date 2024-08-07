package models

type ErrClient struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
