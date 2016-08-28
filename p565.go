package main

import "fmt"
import "math"

// Get number of divisors (cache results)
var numDivisors = make(map[uint64]uint64)

func numDiv(num uint64) uint64 {
	if val, ok := numDivisors[num]; ok {
		return val
	}

	count := uint64(2)
	limit := uint64(math.Sqrt(float64(num)))
	for num%count != 0 {
		// Is number prime?
		if count >= limit {
			numDivisors[num] = num + 1
			return numDivisors[num]
		}
		count += 1
	}

	divisors := numDiv(num / count)

	numDivisors[num] = divisors
	return numDivisors[num]
}

func omDiv(maxNum uint64, divider uint64) uint64 {
	results := uint64(0)

	tick := maxNum / 1000
	for i := uint64(1); i <= maxNum; i++ {
		if i%tick == 0 {
			fmt.Println(float64(float64(i)/(float64(maxNum)/100)), "%")
		}

		if numDiv(i)%divider == 0 {
			results += i
		}
	}

	return results
}

func main() {
	numDivisors[0] = 1
	numDivisors[1] = 1
	numDivisors[2] = 3

	fmt.Println(omDiv(1000000, 2017))
}
