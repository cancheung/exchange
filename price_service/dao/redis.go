package dao

import (
	cf "digicon/price_service/conf"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type RedisCli struct {
	rcon *redis.Client
}

func NewRedisCli() *RedisCli {

	addr := cf.Cfg.MustValue("redis", "addr")
	pass := cf.Cfg.MustValue("redis", "pass")
	num := cf.Cfg.MustInt("redis", "num")
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass, // no password set
		DB:       num,  // use default DB
	})

	pong, err := client.Ping().Result()
	if err != nil {
		log.Fatalf("redis connect faild ")
	}
	log.Infoln(pong)
	return &RedisCli{
		rcon: client,
	}
}

func (s *Dao) GetRedisConn() *redis.Client {
	return s.redis.rcon
}
