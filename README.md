# zapencoder

封装`Golang`内置类型的`Slice`、`Map`结构，以便能够使用`zap.Object()`函数，防止`zap`使用反射进行序列化。

# 使用
```go
import "github.com/sandwich-go/zapgen/zapencoder"
import "go.uber.org/zap"
import "testing"

func TestUint64StringMap(t *testing.T) {
    var a = make(map[uint64]string)
    a[1] = "a"
    a[2] = "b"

    t.Log(zap.Object("A", zapencoder.Uint64StringMap(a)))
}
```