# proxy-pool

[![Actions Status](https://github.com/kagxin/proxy-pool/workflows/Build/badge.svg)](https://github.com/kagxin/proxy-pool/actions)
[![Language](https://img.shields.io/badge/language-gloang-blue.svg)](https://golang.org/)
[![LICENSE](https://img.shields.io/badge/license-MIT-000000.svg)](https://github.com/kagxin/proxy-pool/blob/master/LICENSE)
[![readTheDoc](https://readthedocs.org/projects/golang-proxy-pool/badge/?version=latest)](https://golang-proxy-pool.readthedocs.io)

搭建自己的免费代理池，详细[文档地址](https://golang-proxy-pool.readthedocs.io/)

### 体验地址
* 获取一个proxy [get](http://150.158.87.243:8080/proxy/get)
* 获取所有可用proxy [getall](http://150.158.87.243:8080/proxy/getall)

### 接口描述

|接口|方法|描述|参数|
|-|-|-|-|
|`/get`|GET|随机获取一个可用代理||
|`/getall`|GET|获取所有可用代理||

### 快速开始
* 下载项目的 docker-compose 文件
```bash
docker-compose up
```

* 等待一会儿访问 8080 端口
获取一个proxy `http://localhost:8080/proxy/get`

### 已添加代理源
|代理源|地址|
|-|-|
|全网代理|[地址](http://www.goubanjia.com/)|
|西刺|[地址](https://www.kuaidaili.com/free)|
|IP海|[地址](http://www.ip3366.net/)|
|IP库|[地址](https://ip.jiangxianli.com/)|

> 其他免费代理源欢迎issue提出

### TODO
- [ ] 增加更多免费代理源
- [x] `proxy/get` 接口随机返回可用代理
- [x] log 模块替换
- [x] 添加 readthedocs
- [x] github actions workflow (CI)
