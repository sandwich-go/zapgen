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
}
