# 说明

1、这是学习Go语言的相关源码

2、当前源码使用的是GO_PATH管理

GO_PATH=D:\workspace_go

其D:\workspace_go具体路径如下：

```shell
D:\workspace_go>tree
    ├─bin
    ├─pkg
    │  └─mod
    └─src
        └─go_study_proj
            ├─docs
            │  ├─docker
            │  └─images
            ├─function
            ├─goroutine
            ├─grammar
            ├─interface
            ├─package
            │  └─packagedemo
            ├─reflect
            └─struct
```

由上面可知，go_study_proj是属于src下的一个子模块。

src下可以有多个不同的模块(项目)实现，这样就可能通过GO_PATH实现多项目管理。

> 1、新的包管理方案使用是GO MOD,不需要指定路径，同时还可以指定包依赖版本
>
> 2、每个package(目录)下，只能有一个main函数，否则会报 `main redeclared in this block`异常。

