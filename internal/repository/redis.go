package repository

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisRepository struct {
	client *redis.Client
}

func NewRedisRepository(addr string) *RedisRepository {
	client := redis.NewClient(&redis.Options{Addr: addr})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic("Failed to connect to Redis: " + err.Error())
	}

	return &RedisRepository{client: client}
}

func (r *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()

	if err == redis.Nil {
		return "", nil
	}

	if err != nil {
		return "", err
	}

	return val, nil
}
func (r *RedisRepository) Set(ctx context.Context, key string, value string) error {
	err := r.client.Set(ctx, key, value, 24*time.Hour).Err()
	return err
}
