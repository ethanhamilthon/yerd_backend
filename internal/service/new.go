package service

import (
	"word/internal/service/ask"
	"word/internal/service/user"
)

type DB interface {
	ask.DB
	user.DB
}

type Service struct {
	Ask  ask.AskService
	User user.UserService
}

func New(db DB) *Service {
	return &Service{
		Ask:  *ask.New(db),
		User: *user.New(db),
	}
}
