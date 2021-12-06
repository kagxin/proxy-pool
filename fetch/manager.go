package fetch

import (
	"container/list"
	"context"
	"proxy-pool/internal"
	"proxy-pool/stroage"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
)

type Fetcher func(context.Context) ([]*stroage.ProxyEntity, error)

type Option func(*FetcherManager)

func IntervalOption(interval time.Duration) Option {
	return func(f *FetcherManager) {
		f.interval = interval
	}
}

func ConcurrencyOption(concurrency int) Option {
	return func(f *FetcherManager) {
		f.conChan = make(chan struct{}, concurrency)
	}
}

type FetcherManager struct {
	stroage  stroage.Stroage
	fetchers []Fetcher
	interval time.Duration // s
	runState int32

	ctx     context.Context
	cacnel  context.CancelFunc
	conChan chan struct{}
}

func New(stroage stroage.Stroage, oo ...Option) *FetcherManager {
	ctx, cancel := context.WithCancel(context.Background())
	f := &FetcherManager{
		stroage:  stroage,
		interval: internal.Interval,
		ctx:      ctx,
		cacnel:   cancel,
		conChan:  make(chan struct{}, internal.Concurrency),
	}
	for _, o := range oo {
		o(f)
	}
	return f
}

func (fm *FetcherManager) Register(fetchers []Fetcher) {
	fm.fetchers = append(fm.fetchers, fetchers...)
}

func (fm *FetcherManager) Run() {
	go fm.run()
	timeTicker := time.NewTicker(fm.interval)
	for {
		select {
		case <-timeTicker.C:
			go fm.run()
		case <-fm.ctx.Done():
			log.Infof("FetcherManager Run stop!!\n")
			return
		}
	}
}

func (fm *FetcherManager) run() {
	if atomic.LoadInt32(&fm.runState) == internal.Running { // 已经是 runing 状态
		log.Errorln("FetcherManager already run")
		return
	}
	atomic.StoreInt32(&fm.runState, internal.Running)

	log.Infof("FetcherManager fetch begin!!\n")
	var wg sync.WaitGroup
	var proxyList = list.New()            // 并发安全的
	for _, fetcher := range fm.fetchers { // 爬取页面 proxy
		wg.Add(1)
		go func(_fetcher Fetcher) {
			defer func() {
				wg.Done()
				if err := recover(); err != nil {
					log.Errorf("fetcher err: %+v", err)
				}
			}()
			proxys, err := _fetcher(fm.ctx)
			if err != nil {
				log.Errorf("_fetcher(); err:%+v", err)
				return
			}
			for _, proxy := range proxys {
				proxyList.PushBack(proxy)
			}
		}(fetcher)
	}
	wg.Wait()

	// 校验可用性
	for v := proxyList.Front(); v != nil; v = v.Next() {
		fm.conChan <- struct{}{}
		wg.Add(1)
		go func(proxy *stroage.ProxyEntity) {
			defer func() {
				<-fm.conChan
				wg.Done()
			}()
			ok, err := internal.CheckProxyAvailable(fm.ctx, proxy, internal.HttpBinTimeOut)
			if ok && err == nil {
				log.Infof("FetcherManager check [%s] ok", proxy.Proxy)
				if err := fm.stroage.Put(fm.ctx, proxy); err != nil {
					log.Errorf("fm.stroage.Puts(fm.ctx, %+v); err:%+v", proxy, err)
					return
				}
			} else {
				log.Infof("FetcherManager check [%s] faild", proxy.Proxy)
			}
		}(v.Value.(*stroage.ProxyEntity))
	}
	wg.Wait()
	log.Infof("FetcherManager fetch end !!!\n")
}

func (fm *FetcherManager) Stop() {
	defer fm.cacnel()
}
