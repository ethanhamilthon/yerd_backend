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

func (s *AskService) GenerateWord(ID, UserID, UserLanguage, TargetLanguage, Word string, Writer func(string)) error {
	result := ""
	var ResultSteamer = func(SteamText string) {
		Writer(SteamText)
		result = result + SteamText
	}

	//Generate prompt to openai api from user data
	user_prompt, system_prompt := promptGenarate(UserLanguage, TargetLanguage, Word)

	//To testing AskService without connection to openai API
	if !s.isTest {
		s.runOpenaiAPI(user_prompt, system_prompt, ResultSteamer)
	} else {
		s.runMockAPI(user_prompt, ResultSteamer)
	}

	word := entities.Word{
		WordBasic: entities.WordBasic{
			ID:           ID,
			Title:        Word,
			Description:  result,
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
		return errors.New("Error on creating word")
	}

	return nil
}

func (s *AskService) generatePrompt(UserLanguage, TargetLanguage, Word string) (string, error) {
	return "", nil
}

func (s *AskService) runOpenaiAPI(UserPrompt, SystemPrompt string, ResultSteamer func(string)) {
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
		ResultSteamer(response.Choices[0].Delta.Content)
	}
}

// mock function to testing
func (s *AskService) runMockAPI(prompt string, ResultSteamer func(string)) {
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
		ResultSteamer(word + " ")
		time.Sleep(50 * time.Millisecond)
	}
}
