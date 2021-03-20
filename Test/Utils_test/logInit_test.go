/**
 * @Author HatsuneMona
 * @Date  2021-03-09 08:40
 * @Description 测试log模块
 **/
package Utils

import (
	"service/pkg/Utils"
	"testing"
)

func Test_LoggerInit(t *testing.T) {
	t.Run("测试logger初始化情况", func(t *testing.T) {
		Utils.InitLogger()
		defer Utils.Logger.Sync()
		Utils.Logger.Info("LoggerInfo")
		Utils.Logger.Fatal("LoggerFatal")
		Utils.Logger.Error("LoggerError")
	})
}
