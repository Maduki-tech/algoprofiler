package algorithms

func Quick_sort(input []int, low int, high int) []int {
	if low < high {
		pi := partition(input, low, high)

		Quick_sort(input, low, pi-1)
		Quick_sort(input, pi+1, high)
	}
	return input
}

func partition(input []int, low int, high int) int {
	pivot := input[high]

	left := low - 1

	for j := low; j <= high-1; j++ {
		if input[j] < pivot {
			left++
			swap(input, left, j)
		}
	}

	swap(input, left+1, high)
	return left + 1
}

func swap(input []int, low int, high int) {
	input[low], input[high] = input[high], input[low]
}
