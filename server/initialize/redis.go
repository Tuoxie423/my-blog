package initialize

import (
	"os"
	"server/global"

	"github.com/go-redis/redis"
	"go.uber.org/zap"
)

func ConnectRedis() *redis.Client {
	redisCfg := global.Config.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Address,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		global.Log.Error("Redis连接失败:", zap.Error(err))
		os.Exit(1)
	}
	return client
}
