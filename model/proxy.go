package model

import (
	"time"
)

type (
	// Proxy 数据库model定义
	Proxy struct {
		ID        uint64    `gorm:"COLUMN:id; PRIMARY_KEY; AUTO_INCREMENT" json:"id"`
		Schema    string    `gorm:"COLUMN:schema; SIZE:8; DEFAULT:''; NOT NULL" json:"schema"`
		IP        string    `gorm:"COLUMN:ip; TYPE:varchar(64); DEFAULT:''; NOT NULL" json:"ip"`
		Port      int       `gorm:"COLUMN:port; DEFAULT:0; NOT NULL" json:"port"`
		From      int       `gorm:"COLUMN:from; DEFAULT:0; NOT NULL" json:"from"`
		CheckTime time.Time `gorm:"COLUMN:check_time; TYPE:TIMESTAMP; DEFAULT:CURRENT_TIMESTAMP; NOT NULL" json:"check_time"`
		IsDeleted bool      `gorm:"COLUMN:is_deleted; TYPE:TINYINT; DEFAUTL:0; NOT NULL" json:"is_deleted"`
		CTime     time.Time `gorm:"COLUMN:ctime; TYPE:TIMESTAMP; DEFAULT:CURRENT_TIMESTAMP; NOT NULL" json:"ctime"`
		MTime     time.Time `gorm:"COLUMN:mtime; TYPE:TIMESTAMP; DEFAULT:CURRENT_TIMESTAMP; NOT NULL" json:"mtime"`
	}
)

// TableName gorm table name
func (p *Proxy) TableName() string {
	return "proxy"
}

type (
	// IPKuProxy ip库
	IPKuProxy struct {
		Schema string `json:"protocol"`
		IP     string `json:"ip"`
		Port   string `json:"port"`
	}
	// DataBody asf
	DataBody struct {
		NextPageURL string       `json:"next_page_url"`
		Data        []*IPKuProxy `json:"data"`
	}
	// IPKuResponse rsp
	IPKuResponse struct {
		Code int      `json:"code"`
		Msg  string   `json:"msg"`
		Data DataBody `json:"data"`
	}
)
