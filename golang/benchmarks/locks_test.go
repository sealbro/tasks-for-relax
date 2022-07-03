package benchmarks

import "testing"

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseChannel()
	}
}

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseLock()
	}
}
