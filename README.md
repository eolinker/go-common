# aoplatform common

package

```shell
"github.com/eolinker/go-common"
```

控制台通用工具包


## aolabel

定义
```go

type Struct struct{

	Creator auto.Label `json:"creator" aolabel:"user"`

}
```
赋值
```go

v:= &Struct{
Creator: auto.UUID("uuid")
}

list:= make([]*Struct{},0)
...

auto.CompleteLabels(cxt,v)
auto.CompleteLabels(ctx,list)
auto.CompleteLabels(ctx,list...)
```


注册完成器

```go
auto.RegisterService(name, handler)
```