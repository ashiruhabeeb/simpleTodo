package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

const secretKeySize = 292

// JWTMaker is JSON web token maker
type JWTMaker struct {
	accessTokenKey string
}

// NewJWTMaker creates a new JWTMaker
func NewJWTMaker(accessTokenKey string) (TokenMaker, error) {
	if len(accessTokenKey) < secretKeySize {
		return nil, fmt.Errorf("invalid key size: key size must be equal to %d characters", secretKeySize)
	}
	return &JWTMaker{accessTokenKey}, nil
}

// TokenMaker is an interface for managing jwt tokens
type TokenMaker interface{
	// GenerateAccessToken generates a new token 
	GenerateAccessToken(username string, dur time.Duration)(string, *Payload, error)
	// GenerateRefreshToken generates a new token 
	GenerateRefreshToken(username string, dur time.Duration)(string, *Payload, error)
	// VerifyToken checks id the token is valid or otherwise
	VerifyToken(token string) (*Payload, error)
}


// GenerateAccessToken generates a new token 
func (maker *JWTMaker) GenerateAccessToken(username string, dur time.Duration)(string, *Payload, error){
	payload, err := NewPayload(username, dur)
	if err != nil {
		return "", payload, err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenString, err := token.SignedString([]byte(maker.accessTokenKey))
	return tokenString, payload, err
}

// GenerateRefreshToken generates a refresh token
func (maker *JWTMaker) GenerateRefreshToken(username string, dur time.Duration)(string, *Payload, error){
	payload, err := NewPayload(username, dur)
	if err != nil {
		return "", payload, err
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	refreshTokenString, err := refreshToken.SignedString([]byte(maker.accessTokenKey))
	return refreshTokenString, payload, err
}

// VerifyToken checks id the token is valid or otherwise
func (maker *JWTMaker) VerifyToken(token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(maker.accessTokenKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
