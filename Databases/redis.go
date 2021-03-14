/**
 * @Author HatsuneMona
 * @Date  2021-02-19 22:01
 * @Description 初始化 Redis
 **/
package Databases

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var RDB *redis.Client

func init() {
	RDB = redis.NewClient(&redis.Options{
		Addr:     "172.20.0.2:6379",
		Password: "AHzkfP8ViCX&5MPT5x$VDUjeKWHv6uM4^&Q2qfY9FB0qysuo*@JGbyil7Qa%t7mN", //随机生成的，你看到了也没有用
		DB:       0,
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
