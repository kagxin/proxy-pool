# 关于代码结构
### 设计
代理池的代码主体分为3部分
![image](./proxy_pool.svg)
#### HTTP 服务
使用 `gin` 框架做 HTTP 接口服务

#### Check
周期性检查所有代理的可用性使用，检查可用性的原理：通过代理发送请求到 httpbin.org，
检查是否正常响应

#### Fetch
主要功能是从各个提供免费代理的站点，爬取免费的 proxy，然后进行检查后入库

#### Stroage
proxy 数据存储

### 代码目录结构
```bash
.
├── check           # 代理定期校验
│   ├── check.go
│   └── check_test.go
├── docs            # 文档
│   ├── about.md
│   ├── index.md
│   ├── proxy_pool.png
│   └── software.md
├── exmples         # 启动示例
│   └── exmple
│       └── main.go
├── fetch
│   ├── fetcher     # 代理网页爬取
│   │   ├── fetcher_test.go
│   │   ├── ipku.go
│   │   ├── ipyun.go
│   │   ├── request.go
│   │   ├── request_test.go
│   │   └── xichi.go
│   ├── manager.go  # fetcher 并行管理
│   └── manager_test.go
├── http            # get getall api 接口
│   └── api.go
├── internal        # 枚举及校验部分函数
│   ├── meta.go
│   ├── validator.go
│   └── validator_test.go
├── pkg
└── stroage          # stroage interface{} 定义及 内存 stroage 实现
    ├── memory.go
    ├── memory_test.go
    └── stroage.go
├── Dockerfile
├── docker-compose.yaml
├── LICENSE
├── README.md
├── mkdocs.yml
├── go.mod
├── go.sum
```
