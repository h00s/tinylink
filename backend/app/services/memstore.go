package services

import (
	"context"
	"encoding/json"

	"github.com/go-raptor/raptor"
	"github.com/redis/go-redis/v9"
)

type Memstore struct {
	raptor.Service
	client *redis.Client
}

func NewMemstore(c *raptor.Config) *Memstore {
	return &Memstore{
		client: redis.NewClient(&redis.Options{
			Addr:     c.AppConfig["redis_address"].(string),
			Password: c.AppConfig["redis_password"].(string),
			DB:       int(c.AppConfig["redis_db"].(int64)),
		}),
	}
}

func (m *Memstore) Get(key string) (string, error) {
	value := m.client.Get(context.Background(), key)
	if value.Err() == redis.Nil {
		return "", nil
	}

	data, err := value.Result()
	if err != nil {
		return "", err
	}
	return data, nil
}

func (m *Memstore) Set(key string, value interface{}) error {
	data, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return m.client.Set(context.Background(), key, data, 0).Err()
}
