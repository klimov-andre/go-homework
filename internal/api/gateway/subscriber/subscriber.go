package subscriber

import (
	"encoding/json"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"homework/config/storage"
	"homework/internal/storage/models"
)

const channelMovies = "movies"

func NewRedisSubscriber() *Subscriber {
	rdb := redis.NewClient(&redis.Options{
		Addr:	  storage.RedisAddr,
		Password: storage.RedisPassword,
	})

	return &Subscriber{
		rdb: rdb,
		pubSub: rdb.Subscribe(channelMovies),
		channel: make(chan *models.Movie),
		quit: make(chan struct{}),
	}
}

type Subscriber struct {
	rdb *redis.Client
	pubSub *redis.PubSub
	channel chan *models.Movie
	quit chan struct{}
}

func (s *Subscriber) Channel() <-chan *models.Movie {
	return s.channel
}

func (s *Subscriber) Done() {
	s.quit <- struct{}{}
}

func (s *Subscriber) ReceiveMessages() {
	for {
		select {
		case <-s.quit:
			return
		case msg := <-s.pubSub.Channel():
			var m models.Movie
			if err := json.Unmarshal([]byte(msg.Payload), &m); err != nil {
				logrus.Errorf("subscriber unmarshal error: %v", err.Error())
				continue
			}

			s.channel <- &m
		}
	}
}