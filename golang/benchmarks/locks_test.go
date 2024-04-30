package benchmarks

import "testing"

func BenchmarkChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseChannel()
	}
}

func BenchmarkMutexChannel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseMutexChannel()
	}
}

func BenchmarkLock(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseLock()
	}
}

func BenchmarkAtomic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseAtomic()
	}
}
