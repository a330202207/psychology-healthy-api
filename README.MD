# 心理健康

### 目录
```shell
/
├── api
│   └── v1
├── internal
│   ├── cmd
│   ├── consts
│   ├── controller
│   ├── model
│   │   └── entity
│   ├── packed
│   ├── service
│   │   └── internal
│   │       ├── dao
│   │       └── do
│   ├── config    
│   ├── deploy
│   └── docker
└── utility

```

# 代码生成dao

```
gf gen dao -p ./internal/ -c ./manifest/config.yaml -g -t

```

# 配置文件的键名命名方式统一使用小驼峰