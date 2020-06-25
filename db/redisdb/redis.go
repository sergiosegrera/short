package redisdb

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/sergiosegrera/short/db"

	"github.com/sergiosegrera/short/models"
)

type RedisDB struct {
	client *redis.Client
}

// TODO: Full db options
func New(addr string) (db.DB, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}

	return &RedisDB{
		client: client,
	}, err
}

func (r *RedisDB) CreateLink(link *models.Link) error {
	// TODO: Context timeout
	err := r.client.Set(context.Background(), link.Id, link.Url, 0).Err()
	return err
}

func (r *RedisDB) GetLink(id string) (*models.Link, error) {
	// TODO: Context timeout
	url, err := r.client.Get(context.Background(), id).Result()
	if err != nil {
		return nil, err
	}

	return &models.Link{
		Id:  id,
		Url: url,
	}, err
}
