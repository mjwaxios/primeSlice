package prime_dns

func Primes_dns(limit int) []int {
	limit = limit + 1
	// creates a slice of false with length of limit
	bools := make([]bool, limit)
	// implies up to the sqrt of limit
	for k := 2; k*k <= limit; k++ {
		if bools[k] != true {
			for ix := k * k; ix < limit; ix += k {
				bools[ix] = true
			}
		}
	}
	// index of remaining False in bools = a prime number
	primes := []int{}
	for ix := 2; ix < limit; ix++ {
		if bools[ix] != true {
			primes = append(primes, ix)
		}
	}
	return primes
}


