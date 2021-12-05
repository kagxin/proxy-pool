package stroage

import (
	"context"
	"math/rand"
	"sync/atomic"
	"time"

	"golang.org/x/sync/syncmap"
)

var (
	_ Stroage = (*MemoryStroage)(nil)
)

type MemoryStroage struct {
	proxyLen       int64
	proxyContainer syncmap.Map
}

func NewMemoryStroage() *MemoryStroage {
	rand.Seed(time.Now().UnixNano())
	return &MemoryStroage{}
}

func (m *MemoryStroage) Get(ctx context.Context) (entity *ProxyEntity, err error) {
	if m.proxyLen == 0 {
		return nil, ErrNoFound
	}
	randInt := rand.Intn(int(m.proxyLen))
	var i int = 0
	m.proxyContainer.Range(func(key, value interface{}) bool {
		if i == randInt {
			entity = value.(*ProxyEntity)
			return false
		}
		i++
		return true
	})
	return
}

// 添加一个代理 create or update 逻辑
func (m *MemoryStroage) Put(ctx context.Context, proxy *ProxyEntity) (err error) {

	_, loaded := m.proxyContainer.LoadOrStore(proxy.Proxy, proxy)
	if loaded {
		m.proxyContainer.Store(proxy.Proxy, proxy)
	} else {
		atomic.AddInt64(&m.proxyLen, 1)
	}
	return
}

func (m *MemoryStroage) Puts(ctx context.Context, proxys []*ProxyEntity) (err error) {
	for _, proxy := range proxys {
		if err := m.Put(ctx, proxy); err != nil {
			return err
		}
	}
	return
}

func (m *MemoryStroage) GetAll(ctx context.Context) (proxys []*ProxyEntity, err error) {
	m.proxyContainer.Range(func(key, value interface{}) bool {
		proxys = append(proxys, value.(*ProxyEntity))
		return true
	})
	return
}

// proxy: ip:port
func (m *MemoryStroage) Delete(ctx context.Context, proxy string) (err error) {
	_, loaded := m.proxyContainer.LoadAndDelete(proxy)
	if loaded {
		atomic.AddInt64(&m.proxyLen, -1)
	}
	return nil
}
