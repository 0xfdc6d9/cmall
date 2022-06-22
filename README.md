## 项目依赖

本项目采用了一系列 golang 中比较流行的组件来进行开发

- Gin
- Gorm
- mysql
- redis
- godotenv
- jwt-go
- go-mail

使用的 SDK 或调用的 API

- 阿里云 OSS
- 极验
- 支付 FM
- QQ 第三方登录

## 目录结构

```
cmall-go/
├── api
├── cache
├── conf
├── middleware
├── model
├── pkg
│	├── e
│	├── util
│   ├── sdk
├── serializer
├── server
└── service

```

- api：用于定义接口函数

- cache：redis 相关操作

- conf：用于存储配置文件

- middleware：应用中间件

- model：应用数据库模型

- pkg / e：封装错误码

- pkg / util：工具函数

- pkg / sdk: 极验 sdk 核心函数

- serializer：将数据序列化为 json 的函数

- server 路由逻辑处理

- service：接口函数的实现

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env 文件设置环境变量便于使用(建议开发环境使用)

```
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接语句
REDIS_ADDR="127.0.0.1:6379" # Redis端口地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
GIN_MODE="debug"#开发模式下使用debug
OSS_END_POINT="oss-cn-shenzhen.aliyuncs.com"#你的仓库所在的阿里云对象存储地域节点
OSS_ACCESS_KEY_ID=""#阿里云RAM访问控制用户ID
OSS_ACCESS_KEY_SECRET=""#阿里云RAM访问控制KEY
OSS_BUCKET=""#阿里云OSS仓库名
VAILD_EMAIL="http://localhost:8080/#/vaild/email/" #本地环境邮箱验证地址

#本项目用的是163邮箱STMP
SMTP_HOST=""#163是smtp.163.com
SMTP_EMAIL=""#发送邮件的邮箱
SMTP_PASS=""#SMTP服务的通行证

#极验配置
GEETEST_ID=""#极验账号对应的ID（需要申请）
GEETEST_KEY=""#极验账号对应的KEY（需要申请）

#以下是支付FM配置，详情请查阅支付FM文档
FM_Pay_ID=""#支付FM账号对应的ID
FM_Pay_Key=""#支付FM账号对应的KEY
FM_Pay_NotifyURL=""#支付FM回调地址
FM_Pay_ReturnURL=""#支付FM返回地址

#以下是QQ第三方登录配置，详情请查阅QQ开放平台帮助文档
QQ_Client_ID=""#申请QQ登录时的ID
QQ_Client_KEY=""#申请QQ登录时的KEY
QQ_Redirect_URI=""#QQ登录回调地址
QQ_State=""
```

## 运行

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```
git clone https://github.com/congz666/cmall-go.git
cd cmall-go
go run main.go
```

项目运行后启动在 3000 端口
