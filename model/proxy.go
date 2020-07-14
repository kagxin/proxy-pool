package model

import "time"

type (
	// Proxy 数据库model定义
	Proxy struct {
		ID        int       `gorm:"column:id" json:"id"`
		IP        string    `gorm:"column:ip" json:"ip"`
		Port      int       `gorm:"column:port" json:"port"`
		Schema    string    `gorm:"column:schema" json:"schema"`
		From      int       `gorm:"column:from" json:"from"`
		IsDeleted bool      `gorm:"column:is_deleted" json:"is_deleted"`
		CTime     time.Time `gorm:"column:ctime" json:"ctime"`
		MTime     time.Time `gorm:"column:mtime" json:"mtime"`
	}
)

// TableName gorm table name
func (p *Proxy) TableName() string {
	return "proxy"
}
