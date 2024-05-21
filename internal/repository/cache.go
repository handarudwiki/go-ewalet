package repository

import (
	"github.com/allegro/bigcache/v3"
)

type CacheRepository interface {
	Get(key string) ([]byte, error)
	Set(key string, entry []byte) error
}

type cacheRepository struct {
	cache *bigcache.BigCache
}

func NewCache(cacheConnection *bigcache.BigCache) CacheRepository {
	return &cacheRepository{cache: cacheConnection}
}

func (b *cacheRepository) Get(key string) ([]byte, error) {
	return b.cache.Get(key)
}

func (b *cacheRepository) Set(key string, entry []byte) error {
	return b.cache.Set(key, entry)
}
