package _0_二叉查找树

// 必须导入标准库
import "testing"

// 这个称之为测试表，它能够简单的指定测试用例来避免写出重复代码，
// 见  https://github.com/golang/go/wiki/TableDrivenTests
var tests = []struct {
	input int
	output bool
}{
	{6, true},
	{16, false},
	{3, true},
}

func TestSearch(t *testing.T) {
	//    6
	//   /
	//  3
	tree := Node{
		Key: 6,
		Left: &Node{
			Key: 3,
		},
	}

	for i, test := range tests {
		if res := tree.Search(test.input); res != test.output {
			t.Errorf("%d: got %v, expected %v", i, res, test.output)
		}
	}
}

// go test -bench=
func BenchmarkSearch(b *testing.B) {
	tree := &Node{Key: 5}

	for i := 0; i < b.N; i++ {
		tree.Search(6)
	}
}
