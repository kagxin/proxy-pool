package stroage

import (
	"context"
	"errors"
	"time"
)

type (
	ProxyEntity struct {
		Schema     string    `json:"schema"`
		Proxy      string    `json:"string"` // 唯一标示
		Source     string    `json:"source"`
		CheckTime  time.Time `json:"check_time"`
		CreateTime time.Time `json:"create_time"`
	}
)

var (
	ErrNoFound  = errors.New("no found")
	ErrInternal = errors.New("internal error")
)

func IsErrNoFound(err error) (is bool) {
	if err == ErrNoFound {
		return true
	}
	return
}

func IsErrInternal(err error) (is bool) {
	if err == ErrInternal {
		return true
	}
	return
}

type Stroage interface {
	Get(ctx context.Context) (*ProxyEntity, error)
	Put(ctx context.Context, proxy *ProxyEntity) error // 添加一个代理 create or update 逻辑
	Puts(ctx context.Context, proxys []*ProxyEntity) error
	GetAll(ctx context.Context) ([]*ProxyEntity, error)
	Delete(ctx context.Context, proxy string) error // proxy: ip:port
}
