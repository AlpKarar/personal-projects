package cache

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var (
	ErrKeyNotFound = errors.New("key not found")
)

type Cache struct {
	client *redis.Client
}

type Config struct {
	Addr     string
	Password string
	DB       int
}

func NewCache(cfg Config) (*Cache, error) {
	client := redis.NewClient(&redis.Options{
		Addr: cfg.Addr,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &Cache{client: client}, nil
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	val, err := c.client.Get(ctx, key).Result()
	if err == redis.Nil {
		return "", ErrKeyNotFound
	} else if err != nil {
		return "", fmt.Errorf("failed to get key: %w", err)
	}
	return val, nil
}

func (c *Cache) Set(ctx context.Context, key string, value string, expiry int64) error {
	if err := c.client.Set(ctx, key, value, time.Duration(expiry)*time.Second).Err(); err != nil {
		return fmt.Errorf("failed to set key: %w", err)
	}
	return nil
}

func (c *Cache) Close() error {
	return c.client.Close()
}
