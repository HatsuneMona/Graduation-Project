//Package config
//本文件为保存各种config的默认值，在没有正式的conf读取功能前，默认对所有config信息进行赋值。
//@Author      HatsuneMona
//@CreateTime  2021/3/20 13:38
package config

var (
	//MySQLConfig 默认MySQL服务器配置
	MySQLConfig = &mySQL{
		Addr:     "172.20.0.2",
		Port:     3306,
		User:     "hospital",
		Password: "62QppvxZjCGm7c9jdUyQ",
		Database: "hospital",
	}

	//RedisConfig 默认Redis服务器配置
	RedisConfig = &redis{
		Addr:     "172.20.0.2:6379",
		Password: "AHzkfP8ViCX&5MPT5x$VDUjeKWHv6uM4^&Q2qfY9FB0qysuo*@JGbyil7Qa%t7mN",
		Database: 0,
	}
)
