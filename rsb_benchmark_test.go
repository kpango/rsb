package rsb

import (
	"fmt"
	"math"
	"testing"
)

type test struct {
	n int
	k float32
}

func BenchmarkKnuth(b *testing.B) {
	bench(b, knuth)
}

func BenchmarkShufflePartial(b *testing.B) {
	bench(b, shufflePartial)
}

func bench(b *testing.B, f func(int, float32) []Object) {
	tests := make([]test, 0, 10*10)
	var j float32
	for i := 1; i <= 6; i++ {
		pn := int(math.Pow10(i))
		for j = 0.1; j < 1.0; j += 0.1 {
			tests = append(tests, test{
				n: pn,
				k: j,
			})
		}
	}
	b.ReportAllocs()
	b.ResetTimer()
	for _, tt := range tests {
		b.Run(fmt.Sprintf("n:\t%d\nk:\t%f", tt.n, tt.k), func(bb *testing.B) {
			bb.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					f(tt.n, tt.k)
				}
			})
		})
	}
}
