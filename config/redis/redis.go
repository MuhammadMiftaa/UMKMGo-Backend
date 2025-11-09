package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

func (rdb *RedisRepository) Set(ctx context.Context, key, value string, exp time.Duration) error {
	err := rdb.Client.Set(ctx, key, value, exp).Err()
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	return nil
}

func (rdb *RedisRepository) SetNX(ctx context.Context, key, value string, exp time.Duration) (bool, error) {
	res, err := rdb.Client.SetNX(ctx, key, value, exp).Result()
	if err != nil {
		return res, fmt.Errorf("redis: %w", err)
	}

	return res, nil
}

func (rdb *RedisRepository) HSet(ctx context.Context, key string, value map[string]any, exp time.Duration) error {
	err := rdb.Client.HSet(ctx, key, value).Err()
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	err = rdb.Client.Expire(ctx, key, exp).Err()
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	return nil
}

func (rdb *RedisRepository) HGet(ctx context.Context, key, field string) (string, error) {
	value, err := rdb.Client.HGet(ctx, key, field).Result()
	if err != nil {
		return "", fmt.Errorf("redis: %w", err)
	}

	return value, nil
}

func (rdb *RedisRepository) Publish(ctx context.Context, channel, message string) error {
	err := rdb.Client.Publish(ctx, channel, message).Err()
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	return nil
}

func (rdb *RedisRepository) Subscribe(ctx context.Context, channel string) *redis.PubSub {
	pubsub := rdb.Client.Subscribe(ctx, channel)
	return pubsub
}

func (rdb *RedisRepository) Get(ctx context.Context, key string) (string, error) {
	value, err := rdb.Client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("redis: %w", err)
	}

	return value, nil
}

func (rdb *RedisRepository) Del(ctx context.Context, keys ...string) (int64, error) {
	res, err := rdb.Client.Del(ctx, keys...).Result()
	if err != nil {
		return 0, fmt.Errorf("redis: %w", err)
	}

	return res, nil
}

func (rdb *RedisRepository) Exists(ctx context.Context, keys ...string) (int64, error) {
	found, err := rdb.Client.Exists(ctx, keys...).Result()
	if err != nil {
		return 0, fmt.Errorf("redis: %w", err)
	}

	return found, nil
}

func (rdb *RedisRepository) Incr(ctx context.Context, key string) (int64, error) {
	count, err := rdb.Client.Incr(ctx, key).Result()
	if err != nil {
		return 0, fmt.Errorf("redis: %w", err)
	}

	return count, nil
}

func (rdb *RedisRepository) Expire(ctx context.Context, key string, exp time.Duration) error {
	_, err := rdb.Client.Expire(ctx, key, exp).Result()
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	return nil
}

func (rdb *RedisRepository) MGet(ctx context.Context, keys []string) ([]any, error) {
	values, err := rdb.Client.MGet(ctx, keys...).Result()
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}
	if len(values) == 1 && values[0] == nil {
		return nil, nil
	}

	return values, nil
}

func (rdb *RedisRepository) MSet(ctx context.Context, data map[string]any) error {
	_, err := rdb.Client.MSet(ctx, data).Result()
	if err != nil {
		return fmt.Errorf("redis: %w", err)
	}

	return nil
}

func (rdb *RedisRepository) Scan(ctx context.Context, match string, count int64) ([]string, error) {
	var keys []string
	var cursor uint64
	for {
		k, c, err := rdb.Client.Scan(ctx, cursor, match, count).Result()
		if err != nil {
			return nil, fmt.Errorf("redis: %w", err)
		}

		keys = append(keys, k...)
		if c == 0 {
			break
		}

		cursor = c
	}

	return keys, nil
}

func (rdb *RedisRepository) Keys(ctx context.Context, pattern string) ([]string, error) {
	keys, err := rdb.Client.Keys(ctx, pattern).Result()
	if err != nil {
		return nil, fmt.Errorf("redis: %w", err)
	}

	return keys, nil
}
