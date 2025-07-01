# 个人学习go之实践


## Api 接口定义
遵守openapi规范编写接口文档文件。

使用 [gen_api.go](internal/gen_api.go) 基于接口文档生成代码。

### Client 外部服务客户端
方便对接和工作规范，可以基于外部的接口定义编写openapi外部服务接口文档，（推荐使用模型生成-外部接口至少有一个文档）

使用 [gen_api.go](internal/gen_api.go) 基于接口文档生成客户端代码。


## Db 数据库
定义数据库结构初始化SQL文件

使用 [gen_orm.go](internal/gen_orm.go) 基于SQL文件生成orm代码

工作流程：
1. 创建内存数据库 [embedded-postgres](https://github.com/fergusstrange/embedded-postgres)，并执行数据库结构SQL
2. 使用数据库生成框架 [gorm](https://github.com/go-gorm/gorm), 生成对应数据模型和访问接口




### Auth 授权认证


## DI 依赖注入

fx作为依赖注入框架，只关注实例构建和依赖管理，不应被其他代码依赖，fx所有相关的引用都被额外定义在对应包里面的fx.go文件中
基于提供服务的模块实现NewInstance和OnStart、OnStop 用于创建（注入依赖）、初始化、销毁等 

如果移除项目内所有fx文件，手动创建所有实例也可以保证程序运行，不要对fx进行强依赖


## CodeStruct 代码结构

logx
日志实现

config
配置加载和对象

db
数据库连接

server
http服务相关组件


client
外部服务请求客户端

biz
业务包，不同业务在此细分

biz包内：
biz_name 具体业务包，比如用户管理包 user 租赁管理包 rental

biz_name包内：
api http接口实现，提供外部访问服务，非公开
dba 数据库操作包，一切数据库操作在此定义，非公开
dm 数据模型，除了基于gorm生成的基础数据模型 还会在不同业务场景创建更复杂的数据结构，定义于此包。 非公开
service 内部服务提供包, 满足程序内其他业务模块访问包的实现，此包内应实现 biz_name包下定义的接口。非公开



## Config 配置


## Log 日志输出