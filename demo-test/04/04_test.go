package _4

import "testing"

func BenchmarkStringFromAssignment(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAssignment(100)
	}
}

func BenchmarkStringFromAppendJoin(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromAppendJoin(100)
	}
}

func BenchmarkStringFromBuffer(b *testing.B) {
	for n := 0; n < b.N; n++ {
		StringFromBuffer(100)
	}
}

/*

goos: darwin
goarch: amd64
pkg: demo.io/demo-test/04
BenchmarkStringFromAssignment
BenchmarkStringFromAssignment-4   	  204984	      5774 ns/op
BenchmarkStringFromAppendJoin
BenchmarkStringFromAppendJoin-4   	  336711	      3061 ns/op
BenchmarkStringFromBuffer
BenchmarkStringFromBuffer-4       	 1268924	       931 ns/op
PASS
ok  	demo.io/demo-test/04	7.402s

Process finished with exit code 0

*/
