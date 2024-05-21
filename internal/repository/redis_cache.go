package repository

import (
	"context"
	"time"

	"github.com/handarudwiki/golang-ewalet/config"
	"github.com/handarudwiki/golang-ewalet/domain"
	"github.com/redis/go-redis/v9"
)

type redisCachaeRepository struct {
	rdb *redis.Client
}

func NewRedisClient(cnf *config.Config) domain.CacheRepository {
	return &redisCachaeRepository{
		rdb: redis.NewClient(&redis.Options{
			Addr:     cnf.Redis.Addres,
			Password: cnf.Redis.Password,
		}),
	}
}

func (r redisCachaeRepository) Get(key string) ([]byte, error) {
	val, err := r.rdb.Get(context.Background(), key).Result()
	if err != nil {
		return nil, err
	}
	return []byte(val), nil
}

func (r redisCachaeRepository) Set(key string, entry []byte) error {
	return r.rdb.Set(context.Background(), key, entry, 15*time.Minute).Err()
}
