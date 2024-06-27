package play

import (
	"fmt"
	"testing"
	"word/internal/entities"
)

func TestPlay(t *testing.T) {
	expect_strings := getUniqueArray("gotest", 20)
	play := New(&MockDB{})

	//Get all 20 words
	got_words, err := play.GeneratePlay("", 20, "english")
	if err != nil {
		t.Errorf("Error: expected: nil, got: %v", err.Error())
	}
	if len(got_words) != 20 {
		t.Errorf("Error: expected: 20, got: %v", len(got_words))
	}
	for _, got_word := range got_words {
		if ok := containsString(expect_strings, got_word.Title); !ok {
			t.Errorf("Error: expected: true, got: %v", got_word.Title)
		}
	}
	//Get 10 words
	got_words, err = play.GeneratePlay("", 10, "english")
	if err != nil {
		t.Errorf("Error: expected: nil, got: %v", err.Error())
	}
	if len(got_words) != 10 {
		t.Errorf("Error: expected: 20, got: %v", len(got_words))
	}

	//Get 0 words
	got_words, err = play.GeneratePlay("", 0, "english")
	if err != nil {
		t.Errorf("Error: expected: nil, got: %v", err.Error())
	}
	if len(got_words) != 0 {
		t.Errorf("Error: expected: 20, got: %v", len(got_words))
	}

	//Get 100 words
	got_words, err = play.GeneratePlay("", 100, "english")
	if err != nil {
		t.Errorf("Error: expected: nil, got: %v", err.Error())
	}
	if len(got_words) != 20 {
		t.Errorf("Error: expected: 20, got: %v", len(got_words))
	}

	//Get words with not existing language
	got_words, err = play.GeneratePlay("", 20, "german")
	if err == nil {
		t.Errorf("Error: expected: error, got: nil")
	}
}

type MockDB struct {
}

func (db *MockDB) Words(UserID string) ([]entities.Word, error) {
	word_titles := getUniqueArray("gotest", 20)
	words := make([]entities.Word, 0, 20)
	for _, word_title := range word_titles {
		words = append(words, entities.Word{
			WordBasic: entities.WordBasic{
				ToLanguage: "english",
				Title:      word_title,
			},
		})
	}
	return words, nil
}

func getUniqueArray(Word string, Len int) []string {
	array := make([]string, 0, Len)
	for i := 0; i < Len; i++ {
		array = append(array, Word+fmt.Sprint(i))
	}
	return array
}

func containsString(slice []string, item string) bool {
	for _, v := range slice {
		if v == item {
			return true
		}
	}
	return false
}
