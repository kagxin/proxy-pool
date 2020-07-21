# proxy-pool

搭建自己的代理池，可以用爬虫等场景

### 体验地址
* 获取一个proxy [/proxy/get](http://81.68.131.249:9001/proxy/get)
* 获取所有可用proxy [/proxy/getall](http://81.68.131.249:9001/proxy/getall)

### 接口描述

|接口|方法|描述|参数|
|-|-|-|-|
|`/proxy/get`|GET|随机获取一个可用代理||
|`/proxy/getall`|GET|获取所有可用代理||

### 接口字段描述
|字段|类型|描述|
|-|-|-|
|`ip`|string|代理ip地址|
|`port`|int|代理端口|
|`schema`|string|代理类型|
|`last_check_time`|string|上次检查可用性的时间|

```json
{
    "id": 886,
    "ip": "144.217.101.245",
    "port": 3129,
    "schema": "HTTP",
    "last_check_time": "2020-07-21T10:08:41+08:00"
}
```


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
- [ ] 添加gorm支持的其他数据库
