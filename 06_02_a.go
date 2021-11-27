

	package main

	import (
		"fmt"
		"math"
	)

	func premiers_1 (n int) bool {
		for i := 2; i < n; i++ {
			if (n % i) == 0 {
				return false
			}
		}
	return true
	}

	func premiers_2 (n int) bool {
		k := math.Floor(float64(n/2 + 1))
		for i := 2; i < int(k); i++ {
			if (n % i) == 0 {
				return false
			}
		}
		return true
	}

	func main() {
		fmt.Println(premiers_1(5))
		fmt.Println(premiers_1(5))
	}