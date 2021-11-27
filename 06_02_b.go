

	package main

	import (
		"testing"
	)

	var result bool

	func benchmarkP1(b *testing.B, n int) {
		var r bool
		for i := 0; i < b.N; i++ {
			r = premiers_1(n)
		}
		result = r
	}

	func benchmarkP2(b *testing.B, n int) {
		var r bool
		for i := 0; i < b.N; i++ {
			r = premiers_2(n)
		}
		result = r
	}

	func Benchmark1P1(b *testing.B) {
		benchmarkP1(b, 76303)
	}

	func Benchmark1P2(b *testing.B) {
		benchmarkP2(b, 76303 )
	}

	func Benchmark2P1(b *testing.B) {
		benchmarkP1(b, 212999)
	}

	func Benchmark2P2(b *testing.B) {
		benchmarkP2(b, 212999)
	}