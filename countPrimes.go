package main

import (
	"math"
)

func PrimeNumber(n int) int {
	if n < 2 {
		return 0
	}

	isPrime := []bool{}
	for i := 0; i < n; i++ {
		isPrime = append(isPrime, true)
	}

	isPrime[0], isPrime[1] = false, false

	for i := 2; i < int(math.Ceil(math.Sqrt(float64(n)))); i++ {
		if isPrime[i] {
			for j := i * i; j < n; j += i {
				isPrime[j] = false
			}
		}
	}

	count := 0
	for i := 0; i < len(isPrime); i++ {
		if isPrime[i] {
			count++
		}
	}
	return count
}
