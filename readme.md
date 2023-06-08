
# go-admin

一个简单的小项目
用于go 微服务的实践

## Feature

* 注册中心： consul
* 配置中心： nacos
* orm: gorm
* 环境搭建： docker docker-compose
* 权限鉴定： metadata+token
* 请求安全： TLS
* 链路追踪：
* 熔断：
* 日志记录： zap
* 


## quick start
```shell
# 启动consul
make consul
# consul 默认开启了acl， 初始时需要进入容器生成一下token
consul acl bootstrap

# 得到类似：
AccessorID:       f5d599b2-63d4-1388-5bb7-5ae16cd01237
SecretID:         6f91b479-7245-e4e5-14ea-10b58fc2f888
Description:      Bootstrap Token (Global Management)
Local:            false
Create Time:      2023-06-06 00:03:51.776925006 +0000 UTC
Policies:
   00000000-0000-0000-0000-000000000001 - global-management
   
# 保存SecretID, 测试可以直接使用这个token，生成的话还需生成新的token，并配置策略


# 启动grpc 服务
make server
```


## 权限鉴定和请求安全
权限鉴定使用的是meta做简单的判断，请求安全是使用的TLS加密请求数据,需要注意的是这里并没有使用CA证书授权的机制，因为暂时不打算做客户请求公钥的需求，减少一下项目的复杂度。


## licence

MIT