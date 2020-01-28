# 微服务go项目

    go语言项目目录
## 目录结构
当前目录就是go mod的module所在目录，module名称为qianlitech.com/cloud/backend，见go.mod文件     
目录简介：   

    ├── pkg             # 公共go语言包
    │   ├── api         # 例如 api 包
    ├── services        # 具体服务目录
    │   ├── example     # 例如 example 服务
    ├── third_party     # 第三方文件目录

#内容
    
    API网关、注册中心、链路跟踪、grpc、分布式事务、分布式锁