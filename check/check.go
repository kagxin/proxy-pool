package check

import (
	"context"
	"proxy-pool/stroage"
	"sync"
	"time"

	"proxy-pool/internal"

	log "github.com/sirupsen/logrus"
)

// Checker 检查IP可用性
type Checker struct {
	interval time.Duration //
	stroage  stroage.Stroage

	ctx     context.Context
	cancel  context.CancelFunc
	conChan chan struct{}
}

// New
func New(s stroage.Stroage, interval time.Duration, concurrency int) *Checker {
	ctx, cancel := context.WithCancel(context.Background())
	return &Checker{
		interval: interval,
		stroage:  s,
		ctx:      ctx,
		cancel:   cancel,
		conChan:  make(chan struct{}, concurrency),
	}
}

func (c *Checker) Run() {
	timeTicker := time.NewTicker(c.interval)
	for {
		select {
		case <-timeTicker.C:
			c.run()
		case <-c.ctx.Done():
			log.Infof("Checker Run stop!!\n")
			return
		}
	}
}

func (c *Checker) Stop() {
	defer c.cancel()
}

func (c *Checker) run() {
	log.Info("check run!!\n")
	var wg sync.WaitGroup
	proxys, err := c.stroage.GetAll(c.ctx)
	if err != nil {
		log.Errorf("c.stroage.GetAll(context.Background());err:%+v", err)
		return
	}

	for _, proxy := range proxys {
		c.conChan <- struct{}{}
		wg.Add(1)
		go func(proxy *stroage.ProxyEntity) {
			defer func() {
				wg.Done()
				<-c.conChan
			}()
			ok, err := internal.CheckProxyAvailable(proxy)
			if err == nil && ok {
				log.Infof("proxy [%s], check ok", proxy.Proxy)
			} else {
				log.Infof("proxy [%s], check faild", proxy.Proxy)
				if err := c.stroage.Delete(c.ctx, proxy.Proxy); err != nil {
					log.Errorf("c.stroage.Delete(context.Background(), %s); err:%+v", proxy.Proxy, err)
					return
				}
			}
		}(proxy)
	}
	wg.Wait()
	log.Info("check end!!\n")
}
