//Package Utils
//文件描述...
//@Author      HatsuneMona
//@CreateTime  2021/4/5 15:57
package Utils

import (
	"reflect"
	"service/pkg/auth"
	"testing"
)

func Test_Token(t *testing.T) {
	t.Run("测试创建token", func(t *testing.T) {
		uid := 123456
		issue := "HatsuneMona"
		endTime := 1
		token, err := auth.CreateToken(issue, uid, endTime)
		if err != nil {
			t.Fatalf("生成Token失败，错误信息：%v", err)
		} else {
			t.Logf("生成的Token字符串如下：%v", token)
		}
	})
	uid := 123456
	issue := "HatsuneMona"
	endTime := 1
	token, _ := auth.CreateToken(issue, uid, endTime)

	t.Run("测试解析Token", func(t *testing.T) {
		info, err := auth.ParseToken(token)
		if err != nil {
			t.Fatalf("解析Token失败，错误信息：%v", err)
		} else {
			t.Logf("解析到Token的原始信息如下\n%v\n变量类型为：%v", info, reflect.TypeOf(info))
		}
	})

	//这是一个2021.04.05  16:19创建的token
	outOfDateToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTc2MTQzMDYsImlzcyI6IkhhdHN1bmVNb25hIiwidWlkIjoxMjM0NTZ9.1Xl1p8ELvdRvyXX_fgVwEDBsN0WylpuwzRFsScPZVKE"

	t.Run("测试解析过期的Token", func(t *testing.T) {
		info, err := auth.ParseToken(outOfDateToken)
		if err != nil {
			t.Fatalf("解析Token失败，错误信息：%v", err)
		} else {
			t.Logf("解析到Token的原始信息如下\n%v\n变量类型为：%v", info, reflect.TypeOf(info))
		}
	})
}
