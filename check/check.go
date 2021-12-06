package check

import (
	"context"
	"proxy-pool/stroage"
	"sync"
	"sync/atomic"
	"time"

	"proxy-pool/internal"

	log "github.com/sirupsen/logrus"
)

type Option func(*Checker)

func IntervalOption(interval time.Duration) Option {
	return func(c *Checker) {
		c.interval = interval
	}
}

func ConcurrencyOption(concurrency int) Option {
	return func(c *Checker) {
		c.conChan = make(chan struct{}, concurrency)
	}
}

// Checker 检查IP可用性
type Checker struct {
	interval time.Duration // 检查时间间隔
	stroage  stroage.Stroage
	runState int32

	ctx     context.Context
	cancel  context.CancelFunc
	conChan chan struct{}
}

// New
func New(s stroage.Stroage, oo ...Option) *Checker {

	ctx, cancel := context.WithCancel(context.Background())
	checker := &Checker{
		interval: internal.Interval,
		stroage:  s,
		ctx:      ctx,
		cancel:   cancel,
		conChan:  make(chan struct{}, internal.Concurrency),
	}
	for _, o := range oo {
		o(checker)
	}
	return checker
}

func (c *Checker) Run() {
	timeTicker := time.NewTicker(c.interval)
	for {
		select {
		case <-timeTicker.C:
			go c.run()
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
	if atomic.LoadInt32(&c.runState) == internal.Running {
		log.Errorln("Checker already run")
		return
	}
	atomic.StoreInt32(&c.runState, internal.Running)
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
				<-c.conChan
				wg.Done()
			}()
			ok, err := internal.CheckProxyAvailable(c.ctx, proxy, internal.HttpBinTimeOut)
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
