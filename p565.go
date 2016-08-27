package main

import "fmt"
import "math"

// Get number of divisors (cache results)
var numDivisors = make(map[uint64]*[]uint64)

func remDupes(slice []uint64) *[]uint64 {
	ex := make(map[uint64]bool)

	result := []uint64{}
	for i := range slice {
		if !ex[slice[i]] {
			ex[slice[i]] = true
			result = append(result, slice[i])
		}
	}

	return &result
}

func arrMap(arr []uint64, f func(uint64) uint64) *[]uint64 {
	result := make([]uint64, len(arr))

	for i, val := range arr {
		result[i] = f(val)
	}

	return &result
}

func numDiv(num uint64) *[]uint64 {
	if val, ok := numDivisors[num]; ok {
		return val
	}

	count := uint64(2)
	limit := uint64(math.Sqrt(float64(num)))
	for num%count != 0 {
		// Is number prime?
		if count >= limit {
			numDivisors[num] = &[]uint64{num, 1}
			return numDivisors[num]
		}
		count += 1
	}

	divisors := numDiv(num / count)
	mulDivisors := arrMap(*divisors, func(x uint64) uint64 {
		return x * count
	})
	divisors = remDupes(append(*divisors, *mulDivisors...))

	numDivisors[num] = divisors
	return numDivisors[num]
}

func arrSum(array []uint64) uint64 {
	sum := uint64(0)

	for i := 0; i < len(array); i++ {
		sum += array[uint64(i)]
	}

	return sum
}

func omega(num uint64) uint64 {
	return arrSum(*numDiv(num))
}

func omDiv(maxNum uint64, divider uint64) uint64 {
	results := []uint64{}

	for i := uint64(1); i <= maxNum; i++ {
		if i%(maxNum/1000) == 0 {
			fmt.Println(float64(float64(i)/(float64(maxNum)/100)), "%")
		}

		if omega(i)%divider == 0 {
			results = append(results, i)
		}
	}

	return arrSum(results)
}

func main() {
	numDivisors[0] = &[]uint64{1}
	numDivisors[1] = &[]uint64{1}
	numDivisors[2] = &[]uint64{1, 2}

	fmt.Println(omDiv(1000000, 2017))
}
