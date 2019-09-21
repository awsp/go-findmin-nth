package main

import "fmt"

func FindBoundary(input []int) (int, int) {
	var inputMax = -1
	var inputMin = -1

	// Calculate Min and Max of the input array
	for _, i := range input {
		if inputMax < i {
			inputMax = i
		}
		if inputMin == -1 || i < inputMin {
			inputMin = i
		}
	}

	return inputMax, inputMin
}

func Min(input []int, m int) int {
	if m > len(input) || m <= 0 {
		return -1
	}

	inputMax, inputMin := FindBoundary(input)

	if inputMax > -1 {
		var tmp = make([]int, inputMax - inputMin + 1)
		var output = make([]int, len(input))
		var j = 0

		for _, i := range input {
			tmp[i - inputMin] = i
		}

		for _, t := range tmp {
			if t != 0 {
				output[j] = t
				j++
			}
		}
		return output[m-1]
	}

	return -1
}

func main() {
	var input = []int{
		3, 12, 5, 10, 2, 15, 18,
	}

	fmt.Println(Min(input, 3)) // 10
}
