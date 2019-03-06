package main

import (
	"math"
	"fmt"
	"strconv"
)

func main() {
	primes := make([]int64, 0)
	for i := int64(1); len(primes) < 1000; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}
	//fmt.Println(primes)

	for _, i := range primes {
		for _, j := range primes {
			if j == i {
				continue
			}
			for _, k := range primes {
				if k == i || k == j {
					continue
				}
				for _, l := range primes {
					if l == i || l == j || l == k {
						continue
					}
					for _, m := range primes {
						if m == i || m == j || m == k || m == l {
							continue
						}
						fmt.Println("checking: ", []int64{i, j, k, l, m})
						if checkSet([]int64{i, j, k, l, m}) {
							fmt.Println([]int64{i, j, k, l, m})
							return
						}
					}
				}
			}
		}
	}
}

func checkSet(set []int64) bool {
	for _, i := range set {
		for _, j := range set {
			if i != j && !checkPrimes(i, j) {
				return false
			}
		}
	}
	return true
}

func checkPrimes(n, m int64) bool {
	nm, _ := strconv.Atoi(fmt.Sprintf("%d%d", n, m))
	mn, _ := strconv.Atoi(fmt.Sprintf("%d%d", m, n))

	if isPrime(int64(nm)) && isPrime(int64(mn)) {
		return true
	}

	return false

}

func isPrime(i int64) bool {
	s := math.Sqrt(float64(i))
	for j := int64(2); float64(j) <= s; j++ {
		if i%j == 0 {
			return false
		}
	}
	return true
}
