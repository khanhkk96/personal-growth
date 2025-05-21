package handlers

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

var ctx = context.Background()

type RedisClient struct {
	Client *redis.Client
}

func NewRedis() *RedisClient {
	// Initialize Redis client
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", viper.GetString("REDIS_HOST"), viper.GetString("REDIS_PORT")),
		Password: viper.GetString("REDIS_PASSWORD"),
		DB:       0,
	})
	// Test the connection
	pong, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	fmt.Println(pong) // Should print "PONG"

	return &RedisClient{
		Client: client,
	}
}

func (r *RedisClient) DeleteOneToken(prefix string, uid string, token string) error {
	key := fmt.Sprintf("%s_%s_%s", prefix, uid, token[len(token)-6:])
	// Delete the key
	err := r.Client.Del(ctx, key).Err()
	if err != nil {
		return fmt.Errorf("could not delete key %s: %v", key, err)
	}
	return nil
}

func (r *RedisClient) DeleteOneDevice(uid string, token string) error {
	ackey := fmt.Sprintf("actoken_%s_%s", uid, token[len(token)-6:])
	rfkey := fmt.Sprintf("rftoken_%s_%s", uid, token[len(token)-6:])
	// Delete the key
	err := r.Client.Del(ctx, ackey, rfkey).Err()
	if err != nil {
		return fmt.Errorf("could not delete key %s: %v", []string{ackey, rfkey}, err)
	}
	return nil
}

func (r *RedisClient) DeleteUserToken(uid string) error {
	ackey := fmt.Sprintf("actoken_%s_*", uid)
	// Delete all access tokens for the user
	acIter := r.Client.Scan(ctx, 0, ackey, 0).Iterator()
	for acIter.Next(ctx) {
		log.Println("Val", acIter.Val())
		if err := r.Client.Del(ctx, acIter.Val()).Err(); err != nil {
			fmt.Printf("Failed to delete access token: %s", acIter.Val())
		}
	}

	rfkey := fmt.Sprintf("rftoken_%s_*", uid)
	// Delete all access tokens for the user
	rfIter := r.Client.Scan(ctx, 0, rfkey, 0).Iterator()
	for rfIter.Next(ctx) {
		log.Println("Val", rfIter.Val())
		if err := r.Client.Del(ctx, rfIter.Val()).Err(); err != nil {
			fmt.Printf("Failed to delete refresh token: %s", rfIter.Val())
		}
	}

	return nil
}

func (r *RedisClient) SetVal(key string, val interface{}, ttl time.Duration) error {
	err := r.Client.Set(ctx, key, val, ttl).Err()
	if err != nil {
		return fmt.Errorf("could not set key %s: %s - %v", key, val, err)
	}
	return nil
}

func (r *RedisClient) GetVal(key string) (string, error) {
	val, err := r.Client.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("could not get val with key %s - %v", key, err)
	}
	return val, nil
}
