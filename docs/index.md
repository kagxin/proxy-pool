# proxy-pool

搭建自己的代理池，可以用于爬虫等场景 [github项目地址](https://github.com/kagxin/proxy-pool)

### 体验地址
* 获取一个proxy [/proxy/get](http://150.158.87.243:8080/proxy/get)
* 获取所有可用proxy [/proxy/getall](http://150.158.87.243:8080/proxy/getall)

### 接口描述

|接口|方法|描述|参数|
|-|-|-|-|
|`/proxy/get`|GET|随机获取一个可用代理||
|`/proxy/getall`|GET|获取所有可用代理||

### 接口字段描述
|字段|类型|描述|
|-|-|-|
|`schema`|string|代理类型|
|`proxy`|int|代理端口|
|`source`|string|代理爬取地址|
|`check_time`|string|上次检查可用性的时间|

```json
{
  "schema": "http",
  "proxy": "120.38.32.127:4216",
  "source": "https://ip.jiangxianli.com/api/proxy_ips",
  "check_time": "2021-12-07T02:28:35.110437101Z"
}
```


### 快速开始
* 下载项目的docker-compose文件
```bash
docker-compose up -d
```

* 等待几分钟访问9001端口
获取一个proxy `http://localhost:9001/proxy/get`

