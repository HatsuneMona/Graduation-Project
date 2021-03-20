// @Title       Databases
// @Description 文件描述
// @Author      HatsuneMona
// @CreateTime  2021/3/13 22:24

package Databases

import (
	"context"
	"fmt"
	"service/pkg/cache"
	"testing"
)

func Test_RedisConnection(t *testing.T) {
	t.Run("Redis连接测试", func(t *testing.T) {
		ctx := context.Background()
		str, err := cache.RDB.Ping(ctx).Result()
		if err != nil {
			t.Fatal(fmt.Sprintf("Redis连接错误，错误信息：%v\n", err))
		} else {
			t.Logf("Redis连接成功，ping信息：%v\n", str)
		}
	})
}
