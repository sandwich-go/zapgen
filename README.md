# zapencoder

封装`Golang`内置类型的`Slice`、`Map`结构，以便能够使用`zap.Object()`函数，防止`zap`使用反射进行序列化。

# 使用
```go
var a = make(map[uint64]string)
a[1] = "a"
a[2] = "b"

fmt.Println(zap.Object("A", zapencoder.Uint64StringMap(a)))
```