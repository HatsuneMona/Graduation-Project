//Package Log
//本package包含了log模块的初始化功能，本文件定义了所有error信息。
package errCode

import "encoding/json"

type Errors struct {
	code int
	str  string
}

func (e *Errors) Error() string {
	byteStr, err := json.Marshal(e)
	if err != nil {
		panic(err)
	} else {
		return string(byteStr)
	}
}

//数据库连接错误信息
var (
	MySQLDataBaseConnectError = Errors{
		code: -100001,
		str:  "mysql connect error",
	}
)
