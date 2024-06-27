package service

import (
	"word/internal/service/ask"
	"word/internal/service/play"
	"word/internal/service/user"
	"word/internal/service/word"
)

type DB interface {
	ask.DB
	user.DB
	word.DB
	play.DB
}

type Service struct {
	Ask  *ask.AskService
	User *user.UserService
	Word *word.WordService
	Play *play.PlayService
}

func New(db DB) *Service {
	return &Service{
		Ask:  ask.New(db),
		User: user.New(db),
		Word: word.New(db),
		Play: play.New(db),
	}
}
