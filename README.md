# proxy-pool

[![Actions Status](https://github.com/kagxin/proxy-pool/workflows/Build/badge.svg)](https://github.com/kagxin/proxy-pool/actions)
[![Language](https://img.shields.io/badge/language-gloang-blue.svg)](https://golang.org/)
[![LICENSE](https://img.shields.io/badge/license-MIT-000000.svg)](https://github.com/kagxin/proxy-pool/blob/master/LICENSE)
[![readTheDoc](https://readthedocs.org/projects/golang-proxy-pool/badge/?version=latest)](https://golang-proxy-pool.readthedocs.io)

搭建自己的免费代理池，详细[文档地址](https://golang-proxy-pool.readthedocs.io/)

### 体验地址
* 获取一个proxy [get](http://81.68.131.249:9001/proxy/get)
* 获取所有可用proxy [getall](http://81.68.131.249:9001/proxy/getall)

### 接口描述

|接口|方法|描述|参数|
|-|-|-|-|
|`/get`|GET|随机获取一个可用代理||
|`/getall`|GET|获取所有可用代理||

### 快速开始
* 下载项目的docker-compose文件
```bash
docker-compose up -d
```

* 等待几分钟访问9001端口
获取一个proxy `http://localhost:9001/proxy/get`

### TODO
- [ ] 增加更多免费代理源
- [x] `proxy/get` 接口随机返回可用代理
- [x] log 模块替换
- [x] 添加 readthedocs
- [x] github actions workflow (CI)
