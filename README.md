# zapencoder

封装`Golang`内置类型的`slice`、`map`结构，以便能够使用`zap.Array`或`zap.Object`函数，防止`zap`使用反射进行序列化。

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
	
    var b = make([]int64, 0, 2)
    b = append(b, 1)
    b = append(b, 2)
    t.Log(zap.Array("B", zapencoder.Int64Slice(b)))
}
```