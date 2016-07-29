package luhnmod10

import "testing"

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid("4242424242424242")
	}
}
