package word

import (
	"errors"
	"time"
	"word/internal/entities"
)

type WordService struct {
	db DB
}

type DB interface {
	Words(string) ([]entities.Word, error)
	Languages(string) ([]entities.Language, error)
	CreateWord(entities.Word) error
	Word(string) (entities.Word, error)
	DeleteWord(string, string) error
	UpdateWord(string, string, string, string, time.Time) error
}

func New(db DB) *WordService {
	return &WordService{db}
}

// UserWords returns all words those the user has, and splits them by language
func (s *WordService) UserWords(UserID string) ([]entities.WordsWithLanguage, error) {
	languages, err := s.db.Languages(UserID)
	if err != nil {
		return nil, errors.New("Error get user languages")
	}
	words, err := s.db.Words(UserID)
	if err != nil {
		return nil, errors.New("Error get user words")
	}

	languagedWords := make([]entities.WordsWithLanguage, 0)
	for _, language := range languages {
		words_match_language := make([]entities.Word, 0)
		for _, word := range words {
			if word.ToLanguage == language.LanguageName {
				words_match_language = append(words_match_language, word)
			}
		}
		languagedWords = append(languagedWords, entities.WordsWithLanguage{
			Words:    words_match_language,
			Language: language.LanguageName,
		})
	}

	return languagedWords, nil
}

// CreateManualWord creates a new word in storage
func (s *WordService) CreateManualWord(WordFromUser entities.WordBasic, UserID string) error {
	word := entities.Word{
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    UserID,
		WordBasic: entities.WordBasic{
			ID:           WordFromUser.ID,
			Type:         WordFromUser.Type,
			FromLanguage: WordFromUser.FromLanguage,
			ToLanguage:   WordFromUser.ToLanguage,
			Title:        WordFromUser.Title,
			Description:  WordFromUser.Description,
		},
	}
	err := s.db.CreateWord(word)
	if err != nil {
		return errors.New("Error create word")
	}
	return nil
}

func (s *WordService) Word(ID string) (entities.Word, error) {
	word, err := s.db.Word(ID)
	if err != nil {
		return word, errors.New("Error create word")
	}
	return word, nil
}

func (s *WordService) DeleteWord(ID string, UserID string) error {
	err := s.db.DeleteWord(ID, UserID)
	if err != nil {
		return errors.New("Error delete word")
	}
	return nil
}

func (s *WordService) UpdateWord(ID, Title, Description, UserID string) error {
	now := time.Now()
	err := s.db.UpdateWord(ID, Title, Description, UserID, now)
	if err != nil {
		return errors.New("Error update word")
	}
	return nil
}
