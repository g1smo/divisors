package main

import "fmt"

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

func containsNum(slice []uint64, num uint64) bool {
	for _, el := range slice {
		if el == num {
			return true
		}
	}
	return false
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

	if num == 0 || num == 1 {
		numDivisors[num] = &[]uint64{1}
		return numDivisors[num]
	}

	count := num / 2
	for (count > 1) && (num%count != 0) {
		count -= 1
	}
	factor := num / count

	subDivs := numDiv(count)
	mulDivisors := arrMap(*subDivs, func(num uint64) uint64 {
		return num * factor
	})
	divisors := *remDupes(append(*subDivs, *mulDivisors...))

	numDivisors[num] = &divisors
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
		if i%(maxNum/100) == 0 {
			fmt.Println(i/(maxNum/100), "%")
		}

		if omega(i)%divider == 0 {
			results = append(results, i)
		}
	}

	return arrSum(results)
}

func main() {
	fmt.Println(omDiv(1000000, 2017))
}
