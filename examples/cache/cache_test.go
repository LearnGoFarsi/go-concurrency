package main

import "testing"

func BenchmarkMain(t *testing.B) {
	for i := 0; i < t.N; i++ {
		main()
	}
}
