//Package Controllers
//处理从AdminAPI来的相关请求
//@Author      HatsuneMona
//@CreateTime  2021/4/4 19:10
package Controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"service/Models"
	"service/pkg/auth"
)

func AdminLogin(c *gin.Context) {
	var admin Models.Admin
	admin.Username = c.Param("userName")
	if err := admin.GetAdminByUsername(); err != nil {
		c.JSON(200, msg{
			Code:    -400,
			Message: "用户名或密码错误。",
			Data:    nil,
		})
		return
	}
	if pass, err := auth.PasswordVerify(c.Param("password"), admin.Password); pass == false || err != nil {
		c.JSON(200, msg{
			Code:    -401,
			Message: "用户名或密码错误。",
			Data:    nil,
		})
		return
	}
	//生成Token，传回去
	adminToken, err := auth.CreateToken("MonaHospital", admin.ID, 12)
	if err != nil {
		c.JSON(http.StatusInternalServerError, msg{
			Code:    -500,
			Message: "服务器内部错误，Token创建失败。",
			Data:    nil,
		})
		return
	}
	c.JSON(200, msg{
		Code:    0,
		Message: "登录成功",
		Data: gin.H{
			"token":     adminToken,
			"adminID":   admin.ID,
			"adminName": admin.Name,
		},
	})

}
