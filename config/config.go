package config

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// Config 配置结构体
type Config struct {
	Mysql     *MysqlConfig
	VerifyURL []*VerifyURL
}

// VerifyURL 校验proxy可用性的地址
type VerifyURL struct {
	Schema  string
	TimeOut int
	URL     string
}

// MysqlConfig mysql 配置信息
type MysqlConfig struct {
	Port     int
	Host     string
	Username string
	Password string
	Database string
}

// New 创建一个config
func New() *Config {
	var config = &Config{}
	viper.SetConfigName("conf") // 读取yaml配置文件
	viper.AutomaticEnv()
	confPath := viper.Get("CONF")
	path, ok := confPath.(string)
	if !ok {
		panic(fmt.Sprintf("未找到配置文件:%#v", path))
	}
	viper.AddConfigPath(path)   //设置配置文件的搜索目录
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		log.Fatal(err) // 读取配置文件失败致命错误
	}
	err = viper.UnmarshalKey("prod", config)
	if err != nil {
		panic(err)
	}
	return config
}
