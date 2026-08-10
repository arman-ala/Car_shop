package cache

import (
	"fmt"
	"log"

	"github.com/arman-ala/Car_shop/config"
	redis "github.com/redis/go-redis/v9"
)

var redisClient *redis.Client = nil

func InitRedis(cfg *config.Config) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:            fmt.Sprintf("%v:%v", cfg.Redis.Host, cfg.Redis.Port),
		Username:        cfg.Redis.Host,
		Password:        cfg.Redis.Password,
		DB:              0,
		DialTimeout:     cfg.Redis.DialTimeout,
		ReadTimeout:     cfg.Redis.ReadTimeout,
		WriteTimeout:    cfg.Redis.WriteTimeout,
		PoolTimeout:     cfg.Redis.PoolTimeout,
		PoolSize:        cfg.Redis.PoolSize,
		ConnMaxIdleTime: cfg.Redis.IdleTimeout,
	})
}

func GetRedisClient() *redis.Client {
	if redisClient != nil {
		return redisClient
	}
	log.Println("redis client has not been initialized yet; First create one.")
	return nil
}

func RedisClientClose() {
	if redisClient != nil {
		redisClient.Close()
	}
	log.Println("there is no open redis client.")
}
