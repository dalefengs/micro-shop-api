package redisutils

import (
	"fmt"
	"github.com/go-redis/redis/v8"
	"micro-shop-api/order-web/global"
)

func NewRedisCline() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", global.Config.Redis.Host, global.Config.Redis.Port),
		Password: global.Config.Redis.Password,
	})
}
