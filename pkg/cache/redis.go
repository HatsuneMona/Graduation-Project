/**
 * @Author HatsuneMona
 * @Date  2021-02-19 22:01
 * @Description 初始化 Redis
 **/
package cache

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"service/pkg/config"
)

//RDB 是Redis的连接池
var RDB *redis.Client

func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     config.RedisConfig.Addr,
		Password: config.RedisConfig.Password,
		DB:       config.RedisConfig.Database,
	})
	ctx := context.Background()
	str, err := RDB.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Redis连接错误，错误信息：%v\n", err)
		fmt.Printf("Redis连接错误，Ping信息：%v\n", str)
	} else {
		fmt.Printf("Redis连接成功，ping信息：%v\n", str)
	}
}
