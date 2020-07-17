# proxy-pool

[![Actions Status](https://github.com/kagxin/proxy-pool/workflows/Go/badge.svg)](https://github.com/kagxin/proxy-pool/actions)
[![Language](https://img.shields.io/badge/language-gloang-blue.svg)](https://golang.org/)
[![LICENSE](https://img.shields.io/badge/license-MIT-000000.svg)](https://github.com/kagxin/proxy-pool/blob/master/LICENSE)


爬虫代理池

### 体验地址
* 获取一个proxy [get](http://81.68.131.249:9001/proxy/get)
* 获取所有可用proxy [getall](http://81.68.131.249:9001/proxy/getall)


### docker-compose 启动服务
下载项目的docker-compose文件
```bash
docker-compose up -d
```

### 等待几分钟访问 9001 端口
* 获取一个proxy `http://localhost:9001/proxy/get`
* 获取所有可用proxy `http://localhost:9001/proxy/getall`

### TODO
- [ ] 增加更多免费代理网址
- [x] `proxy/get` 接口随机返回可用代理
- [ ] log 模块替换
- [ ] 添加 readthedocs
- [x] github actions workflow (CI)
