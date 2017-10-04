package auth

import (
	"io/ioutil"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
)

// RSA keys
const (
	publicKeyPath  = "keys/public_key.pub"
	privateKeyPath = "keys/private_key"
)

// SignKey ...
var SignKey []byte

func init() {
	var err error
	SignKey, err = ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatalln("Error reading private key")
		return
	}
}

type jwtClaims struct {
	Name string `json:"name"`
	jwt.StandardClaims
}

// GenerateJWTTokken ...
func GenerateJWTTokken(username string, uuid string) (token string, err error) {
	claims := jwtClaims{
		username,
		jwt.StandardClaims{
			Id:        uuid,
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	}
	rawToken := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)

	token, err = rawToken.SignedString(SignKey)
	if err != nil {
		return "", err
	}

	return token, err
}

// HashPassword ....
func HashPassword(password string) (hash []byte, err error) {
	hash, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return
}

// CompareHash ...
func CompareHash(passwordhash string, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(passwordhash), []byte(password))
	return
}
