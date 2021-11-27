package main

	import (
		"fmt"
		"math"
		"os"
		"runtime/pprof"
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
		cpuFile, err := os.Create("cpuProfile.out")
		if err != nil {
			fmt.Println(err)
			return
		}
		pprof.StartCPUProfile(cpuFile)
		defer pprof.StopCPUProfile()
		total = 0
		for i := 2; i < 100000; i++ {
			n := premiers_2(i)
			if n {
				total = total + 1
			}
		}
		fmt.Println("Total primes:", total)
		fmt.Println()

	}