package utils

import (
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"errors"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"
)

type tokenContextKey struct{}
type Token struct {
	ID  string
	Exp time.Time
}

var (
	TokenKey   = &tokenContextKey{}
	privateKey *ecdsa.PrivateKey
)

func SetupPrivateKey() {
	var err error
	privateKey, err = ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
}

func GenerateToken(id string) (string, error) {
	// Build a new token
	tok, err := jwt.NewBuilder().
		Subject(id).
		Expiration(time.Now().Add(time.Hour * 3)).
		Build()
	if err != nil {
		return "", err
	}

	// Sign the token
	signedToken, err := jwt.Sign(tok, jwt.WithKey(jwa.ES256, privateKey))
	if err != nil {
		return "", err
	}

	return string(signedToken), nil
}

func ValidateToken(ctx context.Context, tokenString string) (context.Context, error) {
	verifiedToken, err := jwt.Parse([]byte(tokenString), jwt.WithKey(jwa.ES256, privateKey))
	if err != nil {
		return ctx, err
	}

	id := verifiedToken.Subject()
	exp := verifiedToken.Expiration()
	if time.Now().After(exp) {
		return ctx, errors.New("token expired")
	}

	token := Token{
		ID:  id,
		Exp: exp,
	}

	ctx = context.WithValue(ctx, TokenKey, token)
	return ctx, nil
}
