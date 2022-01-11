package redis

/*
@author RandySun
@create 2022-01-11-8:35
*/
import (
	"context"
	"fmt"
	"ginWeb/settings"
	"time"

	"github.com/go-redis/redis/v8" // 注意导入的是新版本
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// Init 初始化连接
func Init(cfg *settings.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port,
		),
		Password: cfg.Password, // no password set
		DB:       cfg.DB,       // use default DB
		PoolSize: cfg.PoolSize, //连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 是否连接成功
	_, err = rdb.Ping(ctx).Result()
	return err
}

func Close() {
	_ = rdb.Close()
}
