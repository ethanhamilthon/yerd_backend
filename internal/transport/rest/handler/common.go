package handler

import (
	"fmt"
	"net/http"
	"time"
	"word/config"

	"github.com/golang-jwt/jwt/v5"
)

// JWT GENERATION START
type MyClaims struct {
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
	jwt.RegisteredClaims
}

var jwtKey []byte

func init() {
	jwtKey = []byte(config.JwtKey)
}

func CreateJWT(userID string, userEmail string) (string, error) {
	expirationTime := time.Now().Add(24 * 60 * time.Hour)

	claims := &MyClaims{
		UserID:    userID,
		UserEmail: userEmail,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyJWT(tokenString string) (*MyClaims, error) {
	claims := &MyClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return claims, nil
}

//JWT FUNCS END

// CHECK AUTH START
func CheckAuth(r *http.Request) (string, string, error) {
	token := r.Header.Get("Authorization")
	claims, err := VerifyJWT(token)
	if err != nil {
		fmt.Println(err.Error())
		return "", "", err
	}
	return claims.UserID, claims.UserEmail, nil
}

//CHECK AUTH END

// GET COOKIE
func createCookie(exp time.Time, key string, value string) http.Cookie {
	return http.Cookie{
		Name:     key,
		Value:    value,
		Expires:  exp,
		HttpOnly: false,
		Path:     "/",
	}

}
