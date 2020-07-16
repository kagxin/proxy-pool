package databases

import (
	"fmt"
	"log"
	"proxy-pool/config"
	"proxy-pool/model"

	"github.com/jinzhu/gorm"
)

// DB 结构体
type DB struct {
	Mysql *gorm.DB
}

// New 初始化gorm db
func New(mysql *config.MysqlConfig) *DB {

	var err error
	var db *gorm.DB

	sqlString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		mysql.Username, mysql.Password, mysql.Host, mysql.Port, mysql.Database)
	fmt.Println(sqlString)
	db, err = gorm.Open("mysql", sqlString)
	// TODO: to config file
	db.LogMode(true)
	if err != nil {
		log.Fatalf("mysql connect error %v", err)
	}
	// 创建表结构
	if err := db.AutoMigrate(&model.Proxy{}).Error; err != nil {
		log.Fatalf("mysql create table err:%#v", err)
	}

	return &DB{Mysql: db}
}
