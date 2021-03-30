//Package Controllers
//文件描述...
//@Author      HatsuneMona
//@CreateTime  2021/3/28 16:57
package Controllers

import "github.com/gin-gonic/gin"

func ApiHelloPage(c *gin.Context) {
	c.JSON(200, msg{
		Code:    0,
		Message: "API server start OK",
		Data:    nil,
	})
}
