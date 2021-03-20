//Package password
// @Description 密码加密器
// @Author      HatsuneMona
// @CreateTime  2021/3/14 09:14
package password

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

//@description   明文密码转换为SHA密文密码
//@param         password 		string	明文密码
//@return        hashedPassword string	加密后的SHA密码
//@author        HatsuneMona
//@createTime    2021/3/14 10:33
func PasswordWithSaltGenToSHA(password string) string {
	bytePassword := []byte(password)
	hashedPassword, _ := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	return string(hashedPassword)
}

//@description   将明文密码与SHA密文密码进行匹配验证
//@param         password 		string 	明文密码；
//@param		 hashedPassword	string 	SHA密码
//@return        bool	是否匹配成功；	err	错误信息
//@author        HatsuneMona
//@createTime    2021/3/14 10:34
func PasswordVerify(password string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		return false, errors.New("密码不必配")
	} else if err != nil {
		return false, err
	} else {
		return true, nil
	}
}
