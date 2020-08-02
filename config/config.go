package config

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

// Config 配置结构体
type Config struct {
	Mysql      *MysqlConfig
	VerifyURL  *VerifyURL
	CheckProxy *CheckProxy
	FetchProxy *FetchProxy
	HTTP       *HTTP
}

// HTTP http的配置
type HTTP struct {
	Port string
}

// CheckProxy 检查代理可用性的配置
type CheckProxy struct {
	GoroutineNumber  uint64
	TimeOut          uint64
	CheckAllInterval uint64
}

// FetchProxy 拉取代理的配置
type FetchProxy struct {
	FetchProxyInterval       uint64
	GoroutineNumber          uint64
	FetchSingleProxyInterval uint64
}

// VerifyURL 校验proxy可用性的地址
type VerifyURL struct {
	HTTP  string
	HTTPS string
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
	var mysql = &MysqlConfig{}
	viper.AutomaticEnv()
	mysql.Host = viper.GetString("MYSQL_HOST")
	mysql.Port = viper.GetInt("MYSQL_PORT")
	mysql.Username = viper.GetString("MYSQL_USERNAME")
	mysql.Password = viper.GetString("MYSQL_PASSWORD")
	mysql.Database = viper.GetString("MYSQL_DATABASE")
	config.Mysql = mysql

	confPath := viper.GetString("PROXY_POOL_CONFIG_FILE")
	viper.SetConfigFile(confPath) // 读取yaml配置文件
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // 根据以上配置读取加载配置文件
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败path:%s, err:%#v", confPath, err)) // 读取配置文件失败致命错误
	}
	err = viper.Unmarshal(config)
	if err != nil {
		panic(fmt.Errorf("解析配置文件失败path:%s, err:%#v", confPath, err))
	}
	return config
}
