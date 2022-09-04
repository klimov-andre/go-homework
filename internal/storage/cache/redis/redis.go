package redis

import (
	"context"
	"expvar"
	"fmt"
	"github.com/go-redis/redis"
	"homework/config/storage"
	cachePkg "homework/internal/storage/cache"
	"homework/internal/storage/models"
)

func NewRedisCache(missCounter, hitCounter *expvar.Int) cachePkg.CacheInterface {
	rdb := redis.NewClient(&redis.Options{
		Addr:	  storage.RedisAddr,
		Password: storage.RedisPassword,
	})

	return &cache{
		rdb: rdb,
		missCounter: missCounter,
		hitCounter: hitCounter,
	}
}

// cache implements movie cache for storage via redis
type cache struct {
	rdb *redis.Client

	missCounter *expvar.Int
	hitCounter *expvar.Int
}

func (c *cache) AddOrUpdate(_ context.Context, id uint64, m *models.Movie) {
	if m.Id == 0 {
		m.Id = id
	}
	_ = c.rdb.Set(fmt.Sprint(id), m, 0)
}

func (c *cache) Delete(_ context.Context, id uint64) error {
	return c.rdb.Del(fmt.Sprint(id)).Err()
}

func (c *cache) GetById(_ context.Context, id uint64) (*models.Movie, error) {
	res := c.rdb.Get(fmt.Sprint(id))
	if res.Err() == nil {
		var m models.Movie
		err := res.Scan(&m)
		if err != nil {
			return nil, err
		}
		c.hitCounter.Add(1)
		return &m, nil
	}

	if err := res.Err(); err == redis.Nil {
		c.missCounter.Add(1)
		return nil, cachePkg.ErrCacheNotExists
	} else {
		return nil, err
	}
}