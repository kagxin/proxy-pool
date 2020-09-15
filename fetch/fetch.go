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
	ch      chan *model.Proxy
}

// NewFetcher 新
func NewFetcher(db *databases.DB, conf *config.Config, check *check.Checker) *Fetcher {
	return &Fetcher{
		db:      db,
		conf:    conf,
		checker: check,
	}
}

// CheckAndInsert 检查ip可用性并插入数据库
func (f *Fetcher) CheckAndInsert(proxy *model.Proxy) {
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

}

// FetchAll 拉取所有的代理并检查可用性之后入库
func (f *Fetcher) FetchAll() {
	go f.GetIPYunDaiLi(model.YunDaiLiURL)
	go f.GetIPYunDaiLi(model.YunDaiLiURL2)
	go f.GetQuanWang()
	go f.GetXiChi()
	go f.GetIPKuByAPI()
	go f.GetQiYunProxy()
	go f.Get66Proxy()
}

// GetQuanWang 获取全网代理的免费代理
func (f *Fetcher) GetQuanWang() error {
	_, buf, err := DoRequest(model.QuanWangFetchURL, time.Second*5)
	if err != nil {
		log.Errorf("GetQuanWang DoRequest error:%#v", err)
		return err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		log.Errorf("goquery.NewDocument error:%#v", err)
		return err
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
			return err
		}
		schema := htmlquery.InnerText(htmlquery.FindOne(h, `./td[3]/a/text()`))
		// TODO: from site
		f.CheckAndInsert(&model.Proxy{
			IP:     ipStr,
			Port:   port,
			Schema: schema,
		})
	}
	return nil
}

// GetXiChi 获取西刺的免费代理
func (f *Fetcher) GetXiChi() error {
	_, buf, err := DoRequest(model.KuaiDaiLiFetchURL, time.Second*5)
	if err != nil {
		log.Errorf("XiChiFetchURL DoRequest error:%#v", err)
		return err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		log.Errorf("goquery.NewDocument error:%#v", err)
		return err
	}
	l := htmlquery.Find(doc, `//tbody/tr`)
	for _, h := range l {
		ipStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[1]`))
		portStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[2]]`))
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return err
		}
		schema := htmlquery.InnerText(htmlquery.FindOne(h, `./td[4]`))
		f.CheckAndInsert(&model.Proxy{
			IP:     ipStr,
			Port:   port,
			Schema: schema,
		})
	}

	return nil
}

// GetIPYunDaiLi 获取ip海的免费代理
func (f *Fetcher) GetIPYunDaiLi(url string) error {
	_, buf, err := DoRequest(url, time.Second*5)
	if err != nil {
		log.Errorf("IPSeaURL DoRequest error:%#v", err)
		return err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		log.Errorf("goquery.NewDocument error:%#v", err)
		return err
	}
	l := htmlquery.Find(doc, `//tbody/tr`)
	for _, h := range l {
		ipStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[1]`))
		portStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[2]]`))
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return err
		}
		schema := htmlquery.InnerText(htmlquery.FindOne(h, `./td[4]`))
		f.CheckAndInsert(&model.Proxy{
			IP:     ipStr,
			Port:   port,
			Schema: schema,
		})
	}
	return nil
}

// GetIPKuByAPI 获取ip库
func (f *Fetcher) GetIPKuByAPI() error {
	var APIAddr = model.IPKuURLAPI
	var res model.IPKuResponse
	for true {

		_, buf, err := DoRequest(APIAddr, time.Second*5)
		if err != nil {
			log.Errorf("IPKuURLAPI DoRequest error:%#v", err)
			return err
		}
		err = json.Unmarshal(buf, &res)
		if err != nil {
			return err
		}
		for _, p := range res.Data.Data {
			port, err := strconv.Atoi(p.Port)
			if err != nil {
				continue
			}
			f.CheckAndInsert(&model.Proxy{
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

	return nil
}

// GetQiYunProxy 获取齐云免费proxy
func (f *Fetcher) GetQiYunProxy() error {

	_, buf, err := DoRequest(model.QiYunProxyURL, time.Second*5)
	if err != nil {
		log.Errorf("QiYunProxyURL DoRequest error:%+v", err)
		return err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		log.Errorf("goquery.NewDocument error:%#v", err)
		return err
	}
	l := htmlquery.Find(doc, `//tbody/tr`)
	for _, h := range l {
		ipStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[1]`))
		portStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[2]]`))
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return err
		}
		schema := htmlquery.InnerText(htmlquery.FindOne(h, `./td[4]`))
		f.CheckAndInsert(&model.Proxy{
			IP:     ipStr,
			Port:   port,
			Schema: schema,
		})
	}
	return nil
}

// Get66Proxy 获取66免费proxy
func (f *Fetcher) Get66Proxy() error {
	defer func() {
		if err := recover(); err != nil {
			log.Errorf("Get66Proxy %+v", err)
		}
	}()
	_, buf, err := DoRequest(model.P66ProxyURL, time.Second*5)
	if err != nil {
		log.Errorf("P66ProxyURL DoRequest error:%+v", err)
		return err
	}
	doc, err := htmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		log.Errorf("goquery.NewDocument error:%#v", err)
		return err
	}
	l := htmlquery.Find(doc, `//*[@id="main"]/div/div[1]/table/tbody/tr`)
	for _, h := range l {
		ipStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[1]`))
		portStr := htmlquery.InnerText(htmlquery.FindOne(h, `./td[2]]`))
		port, err := strconv.Atoi(portStr)
		if err != nil {
			continue
		}
		f.CheckAndInsert(&model.Proxy{
			IP:     ipStr,
			Port:   port,
			Schema: "HTTP",
		})
	}
	return nil
}
