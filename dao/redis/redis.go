package redis

import (
	"fmt"
	"ginCli/settings"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.Host, cfg.Port), // Redis服务器地址
		Password: cfg.Password,                             // Redis密码，如果没有设置则留空
		DB:       cfg.DB,
		PoolSize: cfg.PoolSize,
	})

	_, err = rdb.Ping().Result() // 测试连接是否成功
	return
}

func Close() {
	if rdb != nil {
		err := rdb.Close()
		if err != nil {
			zap.L().Error("close Redis failed, err: %v\n", zap.Error(err))
		}
	}
}
