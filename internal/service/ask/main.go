package ask

import (
	"context"
	"errors"
	"io"
	"strings"
	"time"
	"word/config"
	"word/internal/entities"

	"github.com/sashabaranov/go-openai"
)

// Service to create word with OpenAI API
type AskService struct {
	db            DB
	openai_client *openai.Client
	isTest        bool
}

type DB interface {
	CreateWord(entities.Word) error
}

func New(db DB) *AskService {
	client := openai.NewClient(config.OpenaiToken)
	return &AskService{db: db, openai_client: client}
}

func (s *AskService) setTestTrue() {
	s.isTest = true
}

type Writer struct {
	result string
	stream io.Writer
}

// NewWriter returns new writer with field to final result and another writer to stream text
func NewWriter(stream io.Writer) *Writer {
	return &Writer{
		result: "",
		stream: stream,
	}
}

// Write writes current part of text to final result and writes to another writer
func (w *Writer) Write(p []byte) (int, error) {
	w.result = w.result + string(p)
	n, err := w.stream.Write(p)
	return n, err
}

func (s *AskService) GenerateWord(ID, UserID, UserLanguage, TargetLanguage, Word string, Writer io.Writer) error {
	w := NewWriter(Writer)

	//Generate prompt to openai api from user data
	user_prompt, system_prompt := promptGenarate(UserLanguage, TargetLanguage, Word)

	//To testing AskService without connection to openai API
	if !s.isTest {
		s.runOpenaiAPI(user_prompt, system_prompt, w)
	} else {
		s.runMockAPI(user_prompt, w)
	}

	word := entities.Word{
		WordBasic: entities.WordBasic{
			ID:           ID,
			Title:        Word,
			Description:  w.result,
			FromLanguage: UserLanguage,
			ToLanguage:   TargetLanguage,
			Type:         "ai",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    UserID,
	}

	err := s.db.CreateWord(word)
	if err != nil {
		return errors.New("Error create word")
	}

	return nil
}

func (s *AskService) generatePrompt(UserLanguage, TargetLanguage, Word string) (string, error) {
	return "", nil
}

func (s *AskService) runOpenaiAPI(UserPrompt, SystemPrompt string, w io.Writer) {
	ctx := context.Background()
	req := openai.ChatCompletionRequest{
		Model: openai.GPT3Dot5Turbo,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: SystemPrompt,
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: UserPrompt,
			},
		},
		Stream:      true,
		Temperature: 1,
		TopP:        1,
	}
	stream, err := s.openai_client.CreateChatCompletionStream(ctx, req)
	if err != nil {
		return
	}
	defer stream.Close()

	for {
		response, err := stream.Recv()
		if errors.Is(err, io.EOF) {
			break
		}

		if err != nil {
			return
		}

		_, err = w.Write([]byte(response.Choices[0].Delta.Content))
		if err != nil {
			return
		}
	}
}

// mock function to testing
func (s *AskService) runMockAPI(prompt string, w io.Writer) {
	words := strings.Fields(`"Quite" в английском языке означает "довольно", "вполне" или "совсем".
	Это слово используется для усиления прилагательных и наречий, указывая на степень чего-либо.

	1. She is quite talented.
	Она довольно талантлива.
	Здесь "quite" подчеркивает высокий уровень таланта.
	
	2. The movie was quite interesting.
	Фильм был вполне интересным.
	В данном случае "quite" указывает на то, что фильм был действительно интересным, но не чрезмерно.
	
	3. It's quite cold outside.
	На улице довольно холодно.
	Здесь "quite" подчеркивает, что температура заметно низкая.`)
	for _, word := range words {
		_, err := w.Write([]byte(word + " "))
		if err != nil {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}
