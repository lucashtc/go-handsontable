package jwt

import (
	"time"
	"github.com/dgrijalva/jwt-go"
)

var key = []byte("senhaDeAlgumaCoisa")

// Claims ...
type Claims struct {
	UserName string `json:"username"`
	jwt.StandardClaims
}

// GenerateJWT ...
func GenerateJWT(user string) (string, error){
	expirate := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserName : user,
		StandardClaims : jwt.StandardClaims{
			Issuer : "alguma coisa",
			ExpiresAt : expirate.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256 ,claims)
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}