package utils

import (
	"context"
	"errors"
	"time"

	"github.com/lestrrat-go/jwx/v2/jwa"
	"github.com/lestrrat-go/jwx/v2/jwt"

	"github.com/yashikota/chronotes/pkg/redis"
)

type tokenContextKey struct{}
type Token struct {
	UserID  string
	Exp     time.Time
	IsAdmin bool
}

var (
	TokenKey = &tokenContextKey{}
)

func GenerateToken(id string, isAdmin bool) (string, error) {
	// Build a new token
	tok, err := jwt.NewBuilder().
		Subject(id).
		Expiration(time.Now().Add(time.Hour*3)).
		Claim("isAdmin", isAdmin).
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
	isAdminValue, ok := verifiedToken.Get("isAdmin")
	if !ok {
		return ctx, errors.New("isAdmin not found in token")
	}
	isAdmin, ok := isAdminValue.(bool)
	if !ok {
		return ctx, errors.New("isAdmin is not a boolean")
	}

	token := Token{
		UserID:  id,
		Exp:     exp,
		IsAdmin: isAdmin,
	}

	ctx = context.WithValue(ctx, TokenKey, token)
	return ctx, nil
}

func ExtractToken(ctx context.Context) (Token, error) {
	token, ok := ctx.Value(TokenKey).(Token)
	if !ok {
		return Token{}, errors.New("token not found in context")
	}
	return token, nil
}

func GetToken(key string) (string, error) {
	token, err := redis.Client.Get(redis.Ctx, key).Result()
	if err != nil {
		return "", err
	}
	return token, nil
}

func SaveToken(key, token string) error {
	ttl := time.Duration(3) * time.Hour
	err := redis.Client.Set(redis.Ctx, key, token, ttl).Err()
	if err != nil {
		return err
	}
	return nil
}

func DeleteToken(key string) error {
	err := redis.Client.Del(redis.Ctx, key).Err()
	if err != nil {
		return err
	}
	return nil
}
