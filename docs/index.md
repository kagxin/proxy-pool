# proxy-pool

搭建自己的代理池，可以用于爬虫等场景

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

### 使用已有的数据库
> 默认的docker-compose中起了一个mysql做为数据存储的数据库。也可以使用自己已有的数据

* 新建数据库`proxy_pool` `CREATE DATABASE proxy_pool CHARACTER SET utf8mb4`
* 修改docker-compose.yaml
```yaml
version: "3"

services:
  proxy-pool-api:
    image: registry.cn-shanghai.aliyuncs.com/release-lib/proxy-pool:latest
    container_name: proxy-pool-api
    restart: always
    ports:
      - 9001:9001
    volumes:
      - /etc/timezone:/etc/timezone:ro
      - /etc/localtime:/etc/localtime:ro
    environment:
      - PROXY_POOL_CONFIG_FILE=/etc/conf.yaml
      - MYSQL_HOST=your-host
      - MYSQL_PORT=your-port
      - MYSQL_USERNAME=your-username
      - MYSQL_PASSWORD=your-password
      - MYSQL_DATABASE=proxy_pool
    command: api

  proxy-pool-schduler:
    image: registry.cn-shanghai.aliyuncs.com/release-lib/proxy-pool:latest
    container_name: proxy-pool-scheduler
    restart: always
    volumes:
      - /etc/timezone:/etc/timezone:ro
      - /etc/localtime:/etc/localtime:ro
    environment:
      - PROXY_POOL_CONFIG_FILE=/etc/conf.yaml
      - MYSQL_HOST=your-host
      - MYSQL_PORT=your-port
      - MYSQL_USERNAME=your-username
      - MYSQL_PASSWORD=your-password
      - MYSQL_DATABASE=proxy_pool
    command: scheduler
```
