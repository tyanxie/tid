# 简单唯一id生成器

唯一id生成规则：机器ip+进程pid+毫秒时间戳+自增序列

其中机器ip为机器的所有ip地址经过计算得出的hash值

## 下载方式

```shell
go get github.com/tyanxie/tid
```

## 使用方式

```go
// 导包
import "github.com/tyanxie/tid"

// 使用Generate函数
id := tid.Generate()
fmt.Println(id)
```
