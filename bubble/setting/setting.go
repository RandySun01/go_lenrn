package setting

import (
	"gopkg.in/ini.v1"
)

var Conf = new(AppConfig)

// AppConfig 应用程序配置
type AppConfig struct {
	Release      bool          `ini:"release"`    // debug模式是否开启
	Port         int           `ini:"port"`       // 端口
	*MySQLConfig `ini:"mysql"` // 数据配置
}

// MySQLConfig 数据库配置
type MySQLConfig struct {
	User     string `ini:"user"`
	Password string `ini:"password"`
	DB       string `ini:"db"`
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
}

func Init(file string) error {
	return ini.MapTo(Conf, file)
}
