//Package api
//文件描述...
//@Author      HatsuneMona
//@CreateTime  2021/3/28 15:37
package Controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"service/Models"
	"strconv"
)

func GetUserInfoByPhone(c *gin.Context) {
	var user Models.User
	user.Phone = c.Param("phone")
	if err := user.GetUserByPhone(); err != nil {
		c.JSON(400, gin.H{
			"code":    -400,
			"massage": "未找到该用户",
			"err":     fmt.Sprint(err),
		})
		return
	}
	c.JSON(200, msg{
		Code:    0,
		Message: "找到用户",
		Data:    user,
	})
}

func GetUserInfoByID(c *gin.Context) {
	var user Models.User
	var err error
	user.ID, err = strconv.Atoi(c.Param("userID"))
	if err != nil {
		c.JSON(200, msg{
			Code:    -401,
			Message: "用户id非法",
			Data:    nil,
		})
		return
	}
	if err := user.GetUserByPhone(); err != nil {
		c.JSON(400, msg{
			Code:    -400,
			Message: "未找到该用户",
			Data:    nil,
		})
		return
	}
	c.JSON(200, msg{
		Code:    0,
		Message: "找到用户",
		Data:    user,
	})
}
