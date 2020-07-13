package check

import (
	"proxy-pool/config"
	"proxy-pool/databases"
)

// Checker 检查IP可用性
type Checker struct {
	DB   *databases.DB
	Conf *config.Config
}

// NewChecker 检查IP可用性
func NewChecker(db *databases.DB, conf *config.Config) *Checker {
	return &Checker{
		DB:   db,
		Conf: conf,
	}
}

// Check 检查所有IP的可用性
func Check() {

}
