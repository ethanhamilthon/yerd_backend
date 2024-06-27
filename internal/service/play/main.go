package play

import (
	"errors"
	"math/rand"
	"word/internal/entities"
)

type PlayService struct {
	db DB
}

type DB interface {
	Words(string) ([]entities.Word, error)
}

func New(db DB) *PlayService {
	return &PlayService{db}
}

func (s *PlayService) GeneratePlay(UserID string, count int, language string) ([]entities.Word, error) {
	words, err := s.db.Words(UserID)
	if err != nil {
		return nil, errors.New("Error get users")
	}

	//Get all words those has the arg language
	words_match_language := make([]entities.Word, 0)
	for _, word := range words {
		if word.ToLanguage == language {
			words_match_language = append(words_match_language, word)
		}
	}

	if len(words_match_language) == 0 {
		return nil, errors.New("Error there are not matched words")
	} else if len(words_match_language) < count {
		return words_match_language, nil
	}

	return getPlayWords(words_match_language, count), nil
}

func getPlayWords(from []entities.Word, count int) []entities.Word {
	to := make([]entities.Word, 0, count)
	newWords, mediumWords, oldWords := splitSlice(from)

	ratio1, ratio2, ratio3 := 7, 3, 2
	totalRatio := ratio1 + ratio2 + ratio3
	countPart1 := count * ratio1 / totalRatio
	countPart2 := count * ratio2 / totalRatio
	countPart3 := count - countPart1 - countPart2
	selectRandomElements(newWords, &to, countPart1)
	selectRandomElements(mediumWords, &to, countPart2)
	selectRandomElements(oldWords, &to, countPart3)

	return to
}

func selectRandomElements(from []entities.Word, to *[]entities.Word, count int) {
	if count > len(from) {
		count = len(from)
	}

	indices := rand.Perm(len(from))[:count]

	for _, i := range indices {
		*to = append(*to, from[i])
	}

}

func splitSlice[T any](original []T) ([]T, []T, []T) {
	totalLength := len(original)
	ratio1, ratio2, ratio3 := 7, 3, 2
	totalRatio := ratio1 + ratio2 + ratio3

	size1 := totalLength * ratio1 / totalRatio
	size2 := totalLength * ratio2 / totalRatio

	slice1 := original[:size1]
	slice2 := original[size1 : size1+size2]
	slice3 := original[size1+size2:]

	return slice1, slice2, slice3
}
