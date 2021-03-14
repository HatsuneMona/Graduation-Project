/**
 * @Author HatsuneMona
 * @Date  2021-03-09 08:40
 * @Description 测试log模块
 **/
package Utils

import (
	"service/Utils/log"
	"testing"
)

func Test_LoggerInit(t *testing.T) {
	t.Run("测试logger初始化情况", func(t *testing.T) {
		log.InitLogger()
		defer log.Logger.Sync()
		log.Logger.Info("LoggerInfo")
		log.Logger.Fatal("LoggerFatal")
		log.Logger.Error("LoggerError")
	})
}
