package _1

// 单元测试： go test
// 测试覆盖率： go test -cover
// 基准测试： go test -bench=.

import "testing"

func TestTruth(t *testing.T) {
	if true != true {
		t.Fatal("The world is crumbling")
	}
}
