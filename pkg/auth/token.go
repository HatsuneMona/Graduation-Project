//Package auth
//该文件提供了一些token的生成与验证方法。
//@Author      HatsuneMona
//@CreateTime  2021/4/4 19:25
package auth

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var secretKey = []byte("HatsuneMonaIsTheBestOfWorld23333333")

type Claims struct {
	jwt.StandardClaims

	// 追加自己需要的信息
	Uid int `json:"uid"`
	//UserType string `json:"UserType"`
}

//CreateToken
//
//@Description	创建一个新Token
//
//@Param
//				`issuer string`	申请创建者，填写域名
//				`uid int`		生成token的用户uid
//				`timeHour int`	token的过期时间
//
//@Return
//				`tokenString string`	生成的token
//				`err error` 			错误信息
func CreateToken(issuer string, uid int, timeHour int) (tokenString string, err error) {
	c := &Claims{
		jwt.StandardClaims{
			// 过期时间
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(timeHour)).Unix(),
			// 发行者
			Issuer: issuer,
		},
		uid,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err = token.SignedString(secretKey)
	return
}

//ParseToken
//
//@Description	解析Token
//
//@Param
//				`tokenStr string`	tokenString
//
//@Return
//				`claima jwt.Claims`	解析token后得到的结构体
//				`err error`			解析时出现的错误
func ParseToken(tokenStr string) (claims jwt.Claims, err error) {
	var token *jwt.Token
	token, err = jwt.Parse(tokenStr, func(*jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	claims = token.Claims
	return
}
