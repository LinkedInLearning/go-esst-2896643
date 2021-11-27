	package testP

	import (
		"fmt"
		"math"
	)


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
		
		fmt.Println("Nombres premiers:")

	}