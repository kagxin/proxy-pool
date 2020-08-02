package fetch

import (
	"bytes"
	"encoding/json"
	"proxy-pool/check"
	"proxy-pool/config"
	"proxy-pool/databases"
	"proxy-pool/model"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/antchfx/htmlquery"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

// Fetcher 拉取数据
type Fetcher struct {
	db      *databases.DB
	conf    *config.Config
	checker *check.Checker
}

// NewFetcher 新
func NewFetcher(db *databases.DB, conf *config.Config, check *check.Checker) *Fetcher {
	return &Fetcher{
		db:      db,
		conf:    conf,
		checker: check,
	}
}

// FetchAllAndCheck 拉取所有的代理并检查可用性之后入库
func (f *Fetcher) FetchAllAndCheck() {
	var ch = make(chan struct{}, f.conf.FetchProxy.GoroutineNumber)
	var wg sync.WaitGroup
	var allProxys []*model.Proxy
	var proxySites = []func() ([]*model.Proxy, error){
		func() ([]*model.Proxy, error) { return f.GetIPYunDaiLi(model.YunDaiLiURL) },
		func() ([]*model.Proxy, error) { return f.GetIPYunDaiLi(model.YunDaiLiURL2) },
		f.GetQuanWang,
		f.GetXiChi,
		f.GetIPKuByAPI,
	}
	for _, GetFunc := range proxySites {
		proxys, err := GetFunc()
		if err == nil {
			allProxys = append(allProxys, proxys...)
		} else {
			log.Errorf("拉取代理源失败 err:%#v", err)
		}
	}

	for _, proxy := range allProxys {
		ch <- struct{}{}
		wg.Add(1)
		go func(proxy *model.Proxy) {
			defer func() {
				<-ch
				wg.Done()
			}()
			ok, err := f.checker.CheckProxyAvailable(proxy)
			if err != nil || !ok {
				log.Infof("Invalid proxy:%s:%d, %v", proxy.IP, proxy.Port, err)
				return
			}
			log.Infof("Valid proxy:%s:%d.", proxy.IP, proxy.Port)
			// 创建或更新 proxy
			if err := f.db.Mysql.Table("proxy").Where("ip=?", proxy.IP).Where("port=?", proxy.Port).First(&model.Proxy{}).Error; err != nil {
				if gorm.IsRecordNotFoundError(err) {
					if err := f.db.Mysql.Omit("ctime", "mtime", "check_time").Create(proxy).Error; err != nil {
						log.Errorf("f.db.DB.Create ip:%s, port:%d error:%#v", proxy.IP, proxy.Port, err.Error())
					}
				} else {
					log.Errorf("db.DB.Table first %#v", err.Error())
				}
			} else {
				if err := f.db.Mysql.Table("proxy").Where("ip=?", proxy.IP).Where("port=?", proxy.Port).Omit("ctime", "mtime", "check_time").Updates(map[string]interface{}{
					"schema":     proxy.Schema,
					"is_deleted": false,
				}).Error; err != nil {
					log.Errorf("proxy update err:%#v", err.Error())
				}
			}
		}(proxy)
	}
	wg.Wait()
}

// GetQuanWang 获取全网代理的免费代理
func (f *Fetcher) GetQuanWang() ([]*model.Proxy, error) {
	var proxys []*model.Proxy
	_, buf, err := DoRequest(model.QuanWangFetchURL, time.Second*5)
	if err != nil {
		log.Errorf("GetQuanWang DoRequest error:%#v", err)
		return nil, err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		log.Errorf("goquery.NewDocument error:%#v", err)
		return nil, err
	}
	l := htmlquery.Find(doc, `//tbody/tr[@class="success" or @class="warning"]`)
	for _, h := range l {
		ipSplitNode := htmlquery.Find(h, `./td[@class="ip"]/*[not(contains(@style, 'display: none'))
		and not(contains(@style, 'display:none'))
		and not(contains(@class, 'port'))
		]/text()`)
		ipStr := ""
		for _, n := range ipSplitNode {
			ipStr += htmlquery.InnerText(n)
		}
		portStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[@class="ip"]/*[contains(@class, 'port')]`))
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, err
		}
		schema := htmlquery.InnerText(htmlquery.FindOne(h, `./td[3]/a/text()`))
		// TODO: from site
		proxys = append(proxys, &model.Proxy{
			IP:     ipStr,
			Port:   port,
			Schema: schema,
		})
	}

	return proxys, nil
}

// GetXiChi 获取西刺的免费代理
func (f *Fetcher) GetXiChi() ([]*model.Proxy, error) {
	var proxys []*model.Proxy
	_, buf, err := DoRequest(model.KuaiDaiLiFetchURL, time.Second*5)
	if err != nil {
		log.Errorf("XiChiFetchURL DoRequest error:%#v", err)
		return nil, err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		log.Errorf("goquery.NewDocument error:%#v", err)
		return nil, err
	}
	l := htmlquery.Find(doc, `//tbody/tr`)
	for _, h := range l {
		ipStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[1]`))
		portStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[2]]`))
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, err
		}
		schema := htmlquery.InnerText(htmlquery.FindOne(h, `./td[4]`))
		proxys = append(proxys, &model.Proxy{
			IP:     ipStr,
			Port:   port,
			Schema: schema,
		})
	}

	return proxys, nil
}

// GetIPYunDaiLi 获取ip海的免费代理
func (f *Fetcher) GetIPYunDaiLi(url string) ([]*model.Proxy, error) {
	var proxys []*model.Proxy
	_, buf, err := DoRequest(url, time.Second*5)
	if err != nil {
		log.Errorf("IPSeaURL DoRequest error:%#v", err)
		return nil, err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		log.Errorf("goquery.NewDocument error:%#v", err)
		return nil, err
	}
	l := htmlquery.Find(doc, `//tbody/tr`)
	for _, h := range l {
		ipStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[1]`))
		portStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[2]]`))
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return nil, err
		}
		schema := htmlquery.InnerText(htmlquery.FindOne(h, `./td[4]`))
		proxys = append(proxys, &model.Proxy{
			IP:     ipStr,
			Port:   port,
			Schema: schema,
		})
	}

	return proxys, nil
}

// GetIPKuByAPI 获取ip库
func (f *Fetcher) GetIPKuByAPI() ([]*model.Proxy, error) {
	var APIAddr = model.IPKuURLAPI
	var proxys []*model.Proxy
	var res model.IPKuResponse
	for true {

		_, buf, err := DoRequest(APIAddr, time.Second*5)
		if err != nil {
			log.Errorf("IPKuURLAPI DoRequest error:%#v", err)
			return nil, err
		}
		err = json.Unmarshal(buf, &res)
		if err != nil {
			return nil, err
		}
		for _, p := range res.Data.Data {
			port, err := strconv.Atoi(p.Port)
			if err != nil {
				continue
			}
			proxys = append(proxys, &model.Proxy{
				IP:     p.IP,
				Port:   port,
				Schema: strings.ToUpper(p.Schema),
			})

		}
		if res.Data.NextPageURL == APIAddr {
			break
		}
		APIAddr = res.Data.NextPageURL
		time.Sleep(time.Second * time.Duration(f.conf.FetchProxy.FetchSingleProxyInterval))
	}

	return proxys, nil
}
