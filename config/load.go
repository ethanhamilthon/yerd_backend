package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	Mode        string
	Port        string
	OAuthConfig oauth2.Config
	OAuthState  string
	PgConnStr   string
	JwtKey      string
	OpenaiToken string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	Mode = os.Getenv("MODE")
	Port = os.Getenv("PORT")
	OAuthConfig = oauth2.Config{
		ClientID:     os.Getenv("GOOGLEID"),
		ClientSecret: os.Getenv("GOOGLESECRET"),
		RedirectURL:  os.Getenv("GOOGLEREDIRECT"),
		Scopes:       []string{"profile", "email", "https://www.googleapis.com/auth/userinfo.profile", "https://www.googleapis.com/auth/userinfo.email"},
		Endpoint:     google.Endpoint,
	}
	OAuthState = os.Getenv("OAUTH_STATE")
	PgConnStr = os.Getenv("PG_CONN_STR")
	JwtKey = os.Getenv("JWT_KEY")
	OpenaiToken = os.Getenv("OPENAI_TOKEN")

}
