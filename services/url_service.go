package services

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/redis/go-redis/v9"
)

const codeLength = 6
const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const ttl = 24 * time.Hour * 30

type URLService struct {
	rdb *redis.Client
}

func NewURLService(rdb *redis.Client) *URLService {
	return &URLService{rdb: rdb}
}

func (s *URLService) Shorten(originalURL string) (string, error) {
	ctx := context.Background()

	code := generateCode()

	err := s.rdb.Set(ctx, code, originalURL, ttl).Err()
	if err != nil {
		return "", err
	}

	return code, nil
}

func (s *URLService) Resolve(code string) (string, error) {
	ctx := context.Background()

	originalURL, err := s.rdb.Get(ctx, code).Result()
	if errors.Is(err, redis.Nil) {
		return "", errors.New("Not found")
	}
	if err != nil {
		return "", err
	}

	return originalURL, nil
}

func generateCode() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	code := make([]byte, codeLength)
	for i := range code {
		code[i] = charset[r.Intn(len(charset))]
	}
	return string(code)
}
