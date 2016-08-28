package main

import "fmt"
import "math"

// Get number of divisors (cache results)
var numDivisors = make(map[uint64]map[uint64]bool)

func numDiv(num uint64) map[uint64]bool {
	if val, ok := numDivisors[num]; ok {
		return val
	}

	count := uint64(2)
	limit := uint64(math.Sqrt(float64(num)))
	for num%count != 0 {
		// Is number prime?
		if count >= limit {
      numDivisors[num] = map[uint64]bool{num:true, 1:true}
			return numDivisors[num]
		}
		count += 1
	}

	divisors := numDiv(num / count)

	numDivisors[num] = divisors
	return numDivisors[num]
}

func mapSum(numMap map[uint64]bool) uint64 {
	sum := uint64(0)

	for i := range numMap {
		sum += i
	}

	return sum
}

func omega(num uint64) uint64 {
	return mapSum(numDiv(num))
}

func omDiv(maxNum uint64, divider uint64) uint64 {
	results := make(map[uint64]bool)

	for i := uint64(1); i <= maxNum; i++ {
		if i%(maxNum/1000) == 0 {
			fmt.Println(float64(float64(i)/(float64(maxNum)/100)), "%")
		}

		if omega(i)%divider == 0 {
      results[i] = true
		}
	}

	return mapSum(results)
}

func main() {
  numDivisors[0] = map[uint64]bool{1: true}
  numDivisors[1] = map[uint64]bool{1: true}
  numDivisors[2] = map[uint64]bool{1: true, 2: true}

	fmt.Println(omDiv(1000000, 2017))
}
