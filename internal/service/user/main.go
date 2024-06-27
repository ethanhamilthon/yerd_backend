package user

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"
	"word/config"
	"word/internal/entities"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

// Service for google oauth2, get user and update user languages
type UserService struct {
	db DB
}

type DB interface {
	CreateUser(entities.User) error
	UserByEmail(string) (entities.User, error)
	Languages(string) ([]entities.Language, error)
	UpdateUserLanguage(string, string) error
	CreateLanguages([]entities.Language) error
}

func New(db DB) *UserService {
	return &UserService{db}
}

// Google oauth2 login redirect url
func (s *UserService) GoogleLoginURL() string {
	return config.OAuthConfig.AuthCodeURL(config.OAuthState)
}

// Google oauth2 callback processing
// state, code => id, email, err
func (s *UserService) GoogleCallback(state string, code string) (string, string, error) {

	if state != config.OAuthState {
		return "", "", errors.New("State does not match")
	}
	token, err := config.OAuthConfig.Exchange(context.Background(), code)
	if err != nil {
		return "", "", errors.New("Exchang failed")
	}
	userInfo, err := s.getGoogleUserInfo(token)
	if err != nil {
		return "", "", err
	}

	//Parcing fields from interface{} type
	userEmail, ok := userInfo["email"].(string)
	if !ok {
		return "", "", errors.New("Fields parcing failed")
	}
	userAvatar, ok := userInfo["picture"].(string)
	if !ok {
		return "", "", errors.New("Fields parcing failed")
	}
	userFullName, ok := userInfo["name"].(string)
	if !ok {
		return "", "", errors.New("Fields parcing failed")
	}
	userFirstName, ok := userInfo["given_name"].(string)
	if !ok {
		return "", "", errors.New("Fields parcing failed")
	}

	user := entities.User{
		ID:       uuid.New().String(),
		Name:     userFirstName,
		FullName: userFullName,
		Email:    userEmail,
		Avatar:   userAvatar,
		Language: "",
	}

	err = s.db.CreateUser(user)
	if err == nil {
		return user.ID, user.Email, nil
	}

	user, err = s.db.UserByEmail(user.Email)
	if err != nil {
		return "", "", errors.New("Error on get user")
	}

	return user.ID, user.Email, nil
}

// Get info about the user with token by google
func (s *UserService) getGoogleUserInfo(token *oauth2.Token) (map[string]interface{}, error) {
	client := config.OAuthConfig.Client(context.Background(), token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get userinfo: %s", err)
	}
	defer resp.Body.Close()

	var userInfo map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&userInfo)
	if err != nil {
		return nil, fmt.Errorf("failed to parse userinfo response: %s", err)
	}

	return userInfo, nil
}

func (s *UserService) User(UserID string, Email string) (entities.User, []entities.Language, error) {
	user, err := s.db.UserByEmail(Email)
	if err != nil {
		return entities.User{}, nil, errors.New("Error get user")
	}
	languages, err := s.db.Languages(UserID)
	if err != nil {
		return entities.User{}, nil, errors.New("Error get languages")
	}
	return user, languages, nil
}

func (s *UserService) UpdateLanguages(UserLanguage string, TargetLanguages []string, UserID string) error {
	err := s.db.UpdateUserLanguage(UserLanguage, UserID)
	if err != nil {
		return errors.New("Error update user language")
	}

	//Creating new languages
	languages := make([]entities.Language, 0)
	for _, target_language := range TargetLanguages {
		new_language := entities.Language{
			ID:           uuid.New().String(),
			UserID:       UserID,
			LanguageName: target_language,
			CreatedAt:    time.Now(),
		}
		languages = append(languages, new_language)
	}

	err = s.db.CreateLanguages(languages)
	if err != nil {
		return errors.New("Error create target languages")
	}

	return nil
}
