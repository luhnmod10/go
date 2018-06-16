package luhnmod10

import "testing"

func BenchmarkValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Valid("4242424242424242")
	}
}

func BenchmarkChecksum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Checksum("4242424242424242")
	}
}

func BenchmarkCheckDigit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CheckDigit("4242424242424242")
	}
}
