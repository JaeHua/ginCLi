package redis

import (
	"fmt"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
	"github.com/spf13/viper"
)

var rdb *redis.Client

func Init() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")), // Redis服务器地址
		Password: viper.GetString("redis.password"),                                               // Redis密码，如果没有设置则留空
		DB:       viper.GetInt("redis.db"),
		PoolSize: viper.GetInt("redis.pool_size"),
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
