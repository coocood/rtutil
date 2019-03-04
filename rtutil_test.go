package rtutil

import (
	"testing"
)

func BenchmarkAESHash(b *testing.B) {
	buf := make([]byte, 64)
	for i := 0; i < b.N; i++ {
		AESHash(buf)
	}
}

func BenchmarkNanoTime(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NanoTime()
	}
}

func BenchmarkCPUTicks(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CPUTicks()
	}
}

func BenchmarkFastRand(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastRand()
	}
}
