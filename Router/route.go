//Package Router
//本package用于初始化本系统的router
//@Author      HatsuneMona
//@CreateTime  2021/3/20 10:37
package Router

import (
	"github.com/gin-gonic/gin"
	"service/Controllers"
	"service/Router/api"
)

var GinEngine *gin.Engine

func InitRouter() {
	GinEngine = gin.Default()
	apiRouter.Init(GinEngine)
	GinEngine.GET("/", Controllers.ApiHelloPage)
	_ = GinEngine.Run(":8060")
}
