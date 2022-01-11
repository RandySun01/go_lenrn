package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net/http"
)

/*
@author RandySun
@create 2021-12-30-8:56
*/
var customLogger *zap.Logger
var customSugarLogger *zap.SugaredLogger

func main() {
	CustomInitLogger()
	defer customLogger.Sync()
	for {
		simpleHttpGet2("http://www.baidu.com")
		simpleHttpGet2("http://www.baidu.com")
	}

}

// 初始化
func CustomInitLogger() {
	writeSyncer := getLogWriter()
	encoder := getEncoder()
	//  日志格式   写入文件		日志级别
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	customLogger = zap.New(core, zap.AddCaller())
	customSugarLogger = customLogger.Sugar()
}

// 打印日志格式
func getEncoder() zapcore.Encoder {
	// json格式
	//return zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	// 空格分割
	//return zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	// 格式化时间
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 保存文件
//func getLogWriter() zapcore.WriteSyncer {
//	// 每次新建文件写入
//	//file, _ := os.Create("./test.log")
//	// 追加文件
//	file, _ := os.OpenFile("./test.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0744)
//	return zapcore.AddSync(file)
//}

// 保存文件日志切割
func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   "./test.log", // 日志名称
		MaxSize:    1,            // 文件内容大小, MB
		MaxBackups: 5,            // 保留旧文件最大个数
		MaxAge:     30,           // 保留旧文件最大天数
		Compress:   false,        // 文件是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

func simpleHttpGet2(url string) {
	resp, err := http.Get(url)
	if err != nil {
		customSugarLogger.Error(
			"Error fetching url..",
			zap.String("url", url),
			zap.Error(err))
	} else {
		customSugarLogger.Info("Success..",
			zap.String("statusCode", resp.Status),
			zap.String("url", url))
		resp.Body.Close()
	}
}
