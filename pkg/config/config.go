//Package Utils
//本包包含了该项目所有的配置信息。
//@Author      HatsuneMona
//@CreateTime  2021/3/20 13:32
package config

//mySQL
//
//@Description	连接MySQL数据库必要的连接信息
type mySQL struct {
	Addr     string //数据库地址
	Port     int    //数据库端口
	User     string //数据库用户名
	Password string //数据库密码
	Database string //数据库名
}

//redis
//
//@Description	连接redis数据库必要的连接信息
type redis struct {
	Addr     string //数据库地址
	Password string //数据库密码
	Database int    //数据库选择
}

//system
//
//@Description	系统通用配置
type system struct {
}

//net
//
//@Description	网络配置
type net struct {
	Listen string //监听地址
	Port   int    //端口
}
