package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Claims struct {
	UserId int
	jwt.RegisteredClaims
}

//go:generate mockgen -destination=mock_jwt.go -source=jwt.go -package=auth
type TokenAuthentication interface {
	GenerateToken(int) (string, error)
	ValidateToken(string) error
	GetUserIdFromToken(string) int
}

// JwtWrapper wraps the signing key and the issuer
type JwtWrapper struct {
	SecretKey         string
	ExpirationMinutes int64
}

func (jw *JwtWrapper) GenerateToken(Userid int) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Minute * time.Duration(jw.ExpirationMinutes))
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserId: Userid,
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

	claims := Claims{}
	_, err = jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	return
}

func (jw *JwtWrapper) GetUserIdFromToken(tokenString string) (id int) {
	var jwtKey = []byte(jw.SecretKey)

	claims := Claims{}
	_, _ = jwt.ParseWithClaims(
		tokenString,
		&claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		},
	)

	return claims.UserId
}
