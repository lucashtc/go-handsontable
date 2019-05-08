package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
)

var key = []byte("senhaDeAlgumaCoisa")

// Claims ...
type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

// Generate ...ls
func Generate(user string) (string, error) {
	expires := time.Now().Add(24 * time.Hour).Unix()

	claims := &Claims{
		UserName: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expires,
			Issuer:    "user",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", errors.Wrap(err, "falha ao obter Token")
	}
	return tokenString, nil
}

// Refresh ...
func Refresh(tokenString, user string) (string, error) {

	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(j *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if token != nil || token.Valid != true {
		return "", errors.Wrap(err, "Falha ao renovar token")
	}

	tokenString, err = Generate(user)
	if err != nil {
		return "", nil
	}

	return tokenString, nil
}
