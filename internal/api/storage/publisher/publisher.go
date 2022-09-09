package publisher

import (
	"github.com/go-redis/redis"
	"homework/config/storage"
	"homework/internal/storage/models"
)

const channelMovies = "movies"

func NewRedisPublisher() *Publisher {
	rdb := redis.NewClient(&redis.Options{
		Addr:	  storage.RedisAddr,
		Password: storage.RedisPassword,
	})

	return &Publisher{
		rdb: rdb,
	}
}

type Publisher struct {
	rdb *redis.Client
}

func (p *Publisher) Publish(m *models.Movie) {
	res := p.rdb.Publish(channelMovies, m)

	_ = res
}