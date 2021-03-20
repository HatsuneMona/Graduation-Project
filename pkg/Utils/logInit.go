/**
 * @Author HatsuneMona
 * @Date  2021-03-09 08:26
 * @Description 初始化log相关内容
 **/
package Utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var Logger *zap.Logger
var SugarLogger *zap.SugaredLogger

type logInfo struct {
	logFile string //指定日志文件
	debug   bool   //true：输出日志到控制台；false：输出日志到文件
}

var conf logInfo = logInfo{
	logFile: "./test.log",
	debug:   true,
}

func init() {
	InitLogger()
}

func InitLogger() {
	if conf.debug == true {
		Logger = zap.NewExample()
		SugarLogger = Logger.Sugar()
	} else {
		writeSyncer := getLogWriter(conf.logFile)
		encoder := getEncoder()
		core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

		Logger = zap.New(core)
		SugarLogger = Logger.Sugar()
	}
}

func getEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
}

func getLogWriter(filePath string) zapcore.WriteSyncer {
	file, _ := os.Create(filePath)
	return zapcore.AddSync(file)
}
