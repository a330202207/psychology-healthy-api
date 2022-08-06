# 
```shell

/
├── api                 # 对外提供服务的输入/输出数据结构定义 
│   └── v1              # v1版本
├── internal            # 业务逻辑存放目录
│   ├── cmd             # 命令行管理目录
│   ├── consts          # 后台所有常量定义
│   ├── controller      # 接收/解析用户输入参数的入口/接口层
│   ├── dao             # 数据访问对象，这是一层抽象对象，用于和底层数据库交互，仅包含最基础的 CURD 方法
│   │   └── internal    # 业务逻辑存放目录。通过Golang internal特性对外部隐藏可见性
│   ├── logic           # 业务逻辑封装管理，特定的业务逻辑实现和封装
│   ├── model           # 数据结构管理模块，管理数据实体对象，以及输入与输出数据结构定义
│   │   ├── do          # 用于dao数据操作中业务模型与实例模型转换，由工具维护，用户不能修改
│   │   └── entity      # 数据模型是模型与数据集合的一对一关系，由工具维护，用户不能修改
│   └── service         # 用于业务模块解耦的接口定义层。具体的接口实现在logic中进行注入
├── manifest            # 包含程序编译、部署、运行、配置的文件
│   ├── config          # 配置文件存放目录
│   ├── deploy          # 部署相关的文件
│   ├── deploy          # 部署相关的文件
│   ├── docker          # Docker镜像相关依赖文件，脚本文件
│   ├── logs            # 日志文件
│   └── sql             # Sql文件
└── utility             # 公共目录


```

# [代码生成](https://goframe.org/pages/viewpage.action?pageId=1115790)

```
# 代码生成dao
gf gen dao -p ./internal/ -c ./hack/config.yaml -g 

# 生成service
gf gen service

```


# 配置文件的键名命名方式统一使用小驼峰