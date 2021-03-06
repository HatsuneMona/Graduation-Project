module service

go 1.16

require (
	github.com/dgrijalva/jwt-go v3.2.0+incompatible //token 库
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis/v8 v8.7.1
	//latest
	github.com/jinzhu/gorm v1.9.16
	go.uber.org/zap v1.16.0
	golang.org/x/crypto v0.0.0-20210220033148-5ea612d1eb83 //密码SHA相关
	gorm.io/driver/mysql v1.0.4
)
