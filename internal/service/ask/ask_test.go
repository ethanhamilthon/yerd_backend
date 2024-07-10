package ask

import (
	"strings"
	"testing"
	"word/internal/entities"
)

type TestWriter struct {
	count int
}

// NewTestWriter returns a new writer
func NewTestWriter() *TestWriter {
	return &TestWriter{
		count: 0,
	}
}

// Write writes count of write
func (w *TestWriter) Write(p []byte) (int, error) {
	w.count = w.count + 1
	return 0, nil
}

var answer = `"Quite" в английском языке означает "довольно", "вполне" или "совсем".
Это слово используется для усиления прилагательных и наречий, указывая на степень чего-либо.

1. She is quite talented.
Она довольно талантлива.
Здесь "quite" подчеркивает высокий уровень таланта.

2. The movie was quite interesting.
Фильм был вполне интересным.
В данном случае "quite" указывает на то, что фильм был действительно интересным, но не чрезмерно.

3. It's quite cold outside.
На улице довольно холодно.
Здесь "quite" подчеркивает, что температура заметно низкая.`

func TestAsk(t *testing.T) {

	//Setting mocks
	db := &MockDB{}
	service := New(db)
	service.setTestTrue()

	ID := "some_id"
	UserID := "some_user_id"
	UserLanguage := "russian"
	TargetLanguage := "english"
	Word := "quit"

	w := NewTestWriter()

	err := service.GenerateWord(ID, UserID, UserLanguage, TargetLanguage, Word, w)

	if err != nil {
		t.Errorf("Ask: expected: nil, got:%v", err.Error())
	}

	//counter count (count has to be incremented at least 10 times)
	if w.count <= 10 {
		t.Errorf("Ask: expected: >10, got:%v", w.count)
	}

	//Check the datas
	if db.word.ID != ID {
		t.Errorf("Ask: expected: %v, got:%v", ID, db.word.ID)
	}
	if db.word.UserID != UserID {
		t.Errorf("Ask: expected: %v, got:%v", UserID, db.word.UserID)
	}
	if db.word.FromLanguage != UserLanguage {
		t.Errorf("Ask: expected: %v, got:%v", UserLanguage, db.word.FromLanguage)
	}
	if db.word.ToLanguage != TargetLanguage {
		t.Errorf("Ask: expected: %v, got:%v", TargetLanguage, db.word.ToLanguage)
	}
	if db.word.Title != Word {
		t.Errorf("Ask: expected: %v, got:%v", Word, db.word.Title)
	}
	if db.word.Type != "ai" {
		t.Errorf("Ask: expected: %v, got:%v", "ai", db.word.Type)
	}
	if strings.TrimSpace(db.word.Description) != strings.TrimSpace(strings.Join(strings.Fields(answer), " ")+" ") {
		t.Errorf("Ask: expected: %v, got: %v", strings.TrimSpace(strings.Join(strings.Fields(answer), " ")+" "), strings.TrimSpace(db.word.Description))
	}

}

type MockDB struct {
	word entities.Word
}

func (db *MockDB) CreateWord(Word entities.Word) error {
	db.word = Word
	return nil
}

var expected_user_prompt = `Объясните мне, что означает "something" на английском языке (если это другой язык, то переведите на английский). Сначала объясните в общих чертах, что означает это слово/фраза. Затем составьте 3 предложения на английском языке и перевод на русский. И объясните точно в контексте каждого предложения. Ответ нужен без разметки Markdown.`

func TestPromptGenerator(t *testing.T) {
	Word := "something"
	UserLanguage := "russian"
	TargetLanguage := "english"

	user_prompt := promptGenarate(UserLanguage, TargetLanguage, Word)
	if user_prompt != expected_user_prompt {
		t.Errorf("PromptGenarate: expected: %v, got: %v", expected_user_prompt, user_prompt)
	}

}
