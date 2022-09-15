package zapencoder

import (
	"go.uber.org/zap"
	"testing"
)

func TestUint64StringMap(t *testing.T) {
	var a = make(map[uint64]string)
	a[1] = "a"
	a[2] = "b"
	t.Log(zap.Object("map", Uint64StringMap(a)))

	var b = make([]int64, 0, 2)
	b = append(b, 1)
	b = append(b, 2)

	t.Log(zap.Array("B", Int64Slice(b)))
}
