package main

func Sum(numbers []int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	retArraySize := len(numbersToSum)
	ret := make([]int, retArraySize)
	for sliceIndex, sliceArray := range numbersToSum {
		ret[sliceIndex] = Sum(sliceArray)
	}
	return ret
}

func SumAllTails(numbersToSum ...[]int) []int {
	retArraySize := len(numbersToSum)
	ret := make([]int, retArraySize)
	for sliceIndex, sliceArray := range numbersToSum {
		var sum int
		if len(sliceArray) == 0 {
			sum = 0
		} else {
			sum = Sum(sliceArray[1:])
		}
		ret[sliceIndex] = sum
	}
	return ret
}
