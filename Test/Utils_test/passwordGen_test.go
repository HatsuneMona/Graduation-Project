// @Title       Utils
// @Description 文件描述
// @Author      HatsuneMona
// @CreateTime  2021/3/14 09:31

package Utils_test

import (
	"service/pkg/password"
	"testing"
	"time"
)

func Test_PasswordUtil(t *testing.T) {
	t.Run("计算SHA", func(t *testing.T) {
		pw := "passwordTest@1987321654KK"
		t.Logf("明文密码：%v\n", pw)
		sTime := time.Now()
		for i := 0; i < 9; i++ {
			pwSHA := password.PasswordWithSaltGenToSHA(pw)
			t.Logf("第%v次  SHA加密后的密码：%v\n", i+1, string(pwSHA))
		}
		eTime := time.Now()
		t.Logf("9次SHA加密共用时间：%v \t 平均时间：%v", eTime.Sub(sTime), eTime.Sub(sTime)/9)
	})

	pw := "passwordTest@1987321654KK"
	//pwSHA := Utils.PasswordWithSaltGenToSHA(pw)
	t.Logf("明文密码：%v\n", pw)
	pwSHA := make([]string, 9)
	for i := 0; i < 9; i++ {
		pwSHA[i] = password.PasswordWithSaltGenToSHA(pw)
		t.Logf("第%v次  SHA加密后的密码：%v\n", i+1, string(pwSHA[i]))
	}

	t.Run("密码验证（正确向）", func(t *testing.T) {
		sTime := time.Now()
		for i := 0; i < 9; i++ {
			ok, err := password.PasswordVerify(pw, pwSHA[i])
			if ok != true || err != nil {
				t.Fatalf("第%v次  密码验证错误，错误信息：%v", i+1, err)
			} else {
				t.Logf("第%v次  密码对验证成功！", i+1)
			}
		}
		eTime := time.Now()
		t.Logf("9次SHA验证共用时间：%v \t 平均时间：%v", eTime.Sub(sTime), eTime.Sub(sTime)/9)
	})

	wrongPW := "p@$$wordTest@3692581475KK"
	t.Run("密码验证（错误向）", func(t *testing.T) {
		sTime := time.Now()
		for i := 0; i < 9; i++ {
			ok, err := password.PasswordVerify(wrongPW, pwSHA[i])
			if ok != true || err != nil {
				t.Logf("第%v次  密码验证错误，错误信息：%v", i+1, err)
			} else {
				t.Fatalf("第%v次  密码对验证成功！", i+1)
			}
		}
		eTime := time.Now()
		t.Logf("9次SHA验证共用时间：%v \t 平均时间：%v", eTime.Sub(sTime), eTime.Sub(sTime)/9)
	})

	zhPW := "派蒙：前面的区域，以后再来探索吧~"
	t.Run("中文密码计算与验证", func(t *testing.T) {
		t.Logf("明文密码：%v\n", zhPW)
		pwSHA := make([]string, 9)
		for i := 0; i < 9; i++ {
			pwSHA[i] = password.PasswordWithSaltGenToSHA(zhPW)
			t.Logf("第%v次  SHA加密后的密码：%v\n", i+1, pwSHA[i])
		}
		for i := 0; i < 9; i++ {
			ok, err := password.PasswordVerify(zhPW, pwSHA[i])
			if ok != true || err != nil {
				t.Errorf("第%v次  密码验证错误，错误信息：%v", i+1, err)
			} else {
				t.Logf("第%v次  密码对验证成功！", i+1)
			}
		}
	})
}
