## 爬取v2ex文章

> 爬取v2ex文章，有事瞧瞧，没事看看

## Logo

```python
'  ██╗   ██╗██████╗ ███████╗██╗  ██╗     ███████╗██████╗ ██╗██████╗ ███████╗██████╗ 
'  ██║   ██║╚════██╗██╔════╝╚██╗██╔╝     ██╔════╝██╔══██╗██║██╔══██╗██╔════╝██╔══██╗
'  ██║   ██║ █████╔╝█████╗   ╚███╔╝█████╗███████╗██████╔╝██║██║  ██║█████╗  ██████╔╝
'  ╚██╗ ██╔╝██╔═══╝ ██╔══╝   ██╔██╗╚════╝╚════██║██╔═══╝ ██║██║  ██║██╔══╝  ██╔══██╗
'   ╚████╔╝ ███████╗███████╗██╔╝ ██╗     ███████║██║     ██║██████╔╝███████╗██║  ██║
'    ╚═══╝  ╚══════╝╚══════╝╚═╝  ╚═╝     ╚══════╝╚═╝     ╚═╝╚═════╝ ╚══════╝╚═╝  ╚═╝
'                                                                                   
```

## 目标

我们的目标站点是 https://v2ex.com

## 开始

- 分析页面，获取文章列表
- 分析页面，获取描述和评论

### Install
```
$ go get github.com/scloud/v2ex-spider
```

### 运行
```
$ go run main.go
```

### Usage
```
index       获取首页文章列表

index 2     获取首页序号为2的文章
```

