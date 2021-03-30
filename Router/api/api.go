//Package apiRouter
//文件描述...
//@Author      HatsuneMona
//@CreateTime  2021/3/28 15:21
package apiRouter

import (
	"github.com/gin-gonic/gin"
	"service/Controllers"
)

func Init(e *gin.Engine) {
	apiRouter := e.Group("/api")
	{
		userApi := apiRouter.Group("/user")
		{
			userApi.GET("/:phone", Controllers.GetUserInfoByPhone)
		}
	}
}
