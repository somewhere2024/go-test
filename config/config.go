package config

import (
	"github.com/spf13/viper"
	"os"
)

type LogConfig struct {
	Level      string `json:"level"`       // 日志等级
	Filename   string `json:"filename"`    // 基准日志文件名
	MaxSize    int    `json:"maxsize"`     // 单个日志文件最大内容，单位：MB
	MaxAge     int    `json:"max_age"`     // 日志文件保存时间，单位：天
	MaxBackups int    `json:"max_backups"` // 最多保存几个日志文件
}

var Cfg = &LogConfig{
	Level:      "info",
	Filename:   "./logs/app.log",
	MaxSize:    100,
	MaxAge:     30,
	MaxBackups: 10,
}

//下面没用

type DBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	driver   string
}

type Config struct {
	AppName  string
	AppPort  int
	DBConfig DBConfig
}

func ConfigInit() {
	path, _ := os.Getwd()
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")

	// 从配置文件读取配置
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	// 从环境变量读取配置

	viper.SetEnvPrefix("app") // 读取APP_前缀的环境变量
	viper.AutomaticEnv()

}

// 暂时没用

func GetConfig() Config {
	ConfigInit()
	config := Config{
		AppName: viper.GetString("app.name"),
		AppPort: viper.GetInt("app.port"),
		DBConfig: DBConfig{
			Host:     viper.GetString("database.host"),
			Port:     viper.GetInt("database.port"),
			User:     viper.GetString("database.username"),
			Password: viper.GetString("database.password"),
			Name:     viper.GetString("database.name"),
			driver:   viper.GetString("database.driver"),
		},
	}
	return config
}
