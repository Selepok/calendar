package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey         string
	Issuer            string
	ExpirationMinutes int64
}

func (jw *JwtWrapper) GenerateToken(username string) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(jw.ExpirationMinutes))
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	var jwtKey = []byte(jw.SecretKey)

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err = token.SignedString(jwtKey)

	if err != nil {
		return
	}

	return
}

func (jw *JwtWrapper) ValidateToken(tokenString string) (err error) {
	var jwtKey = []byte(jw.SecretKey)

	_, err = jwt.ParseWithClaims(
		tokenString,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	return
}
