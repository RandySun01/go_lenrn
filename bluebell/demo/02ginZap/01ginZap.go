package main

import (
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
@author RandySun
@create 2022-01-04-8:52
*/
var ginLogger *zap.Logger
var ginSugarLogger *zap.SugaredLogger

func main() {
	//gin.Default()
	cfg := &LogConfig{
		Level:      "DEBUG",
		Filename:   "./test.log",
		MaxSize:    1,
		MaxAge:     30,
		MaxBackups: 5,
	}

	GinInitLogger(cfg)
	r := gin.New()
	r.Use(GinLogger(ginLogger), GinRecovery(ginLogger, true))
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello gin zap")

	})

	r.GET("/panic", func(c *gin.Context) {
		panic("gin zap panic")
		c.String(http.StatusOK, "hello gin zap")

	})
	r.Run(":9999")
}

// 初始化
func GinInitLogger(cfg *LogConfig) (err error) {
	writeSyncer := getLogWriter(cfg.Filename, cfg.MaxSize, cfg.MaxBackups, cfg.MaxAge)
	encoder := getEncoder()
	// 日志级别
	var l = new(zapcore.Level)
	err = l.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		return
	}
	//  日志格式   写入文件		日志级别
	core := zapcore.NewCore(encoder, writeSyncer, l)
	ginLogger = zap.New(core, zap.AddCaller())
	zap.ReplaceGlobals(ginLogger) // 替换zap包中全局的logger实例，后续在其他包中只需使用zap.L()调用即可
	ginSugarLogger = ginLogger.Sugar()
	return
}

// 打印日志格式
func getEncoder() zapcore.Encoder {
	// 格式化时间
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 保存文件日志切割
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 日志名称
		MaxSize:    maxSize,   // 文件内容大小, MB
		MaxBackups: maxBackup, // 保留旧文件最大个数
		MaxAge:     maxAge,    // 保留旧文件最大天数
		Compress:   false,     // 文件是否压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

// GinLogger 接收gin框架默认的日志
func GinLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()             // 请求的时间
		path := c.Request.URL.Path      // 请求的时间
		query := c.Request.URL.RawQuery // 请求的参数
		c.Next()                        // 执行后续中间件

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),                                 // 状态码
			zap.String("method", c.Request.Method),                               // 请求的方法
			zap.String("path", path),                                             // 请求的路径
			zap.String("query", query),                                           // 请求的参数
			zap.String("ip", c.ClientIP()),                                       // 请求的IP
			zap.String("user-agent", c.Request.UserAgent()),                      // 请求头
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()), // 错误信息
			zap.Duration("cost", cost),                                           // 请求时间
		)
	}
}

// GinRecovery recover掉项目可能出现的panic stack是否记录堆栈信息
func GinRecovery(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// Check for a broken connection, as it is not really a
				// condition that warrants a panic stack trace.
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					// If the connection is dead, we can't write a status to it.
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				}

				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
				}
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()
		c.Next()
	}
}
