//Package Middlewares
//有关Admin用户的中间件
//@Author      HatsuneMona
//@CreateTime  2021/4/5 16:41
package Middlewares

import (
	"github.com/gin-gonic/gin"
)

//LoginAuthBasics
//
//@Description	Admin用户验证的基础验证，只验证是否登录过
//
//@Return
//				gin.HandlerFunc
func LoginAuthBasics() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
