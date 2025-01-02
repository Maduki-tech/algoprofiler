package algorithms

func Heap_sort(input []int) []int {
	n := len(input)

	// Building the heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(input, n, i)
	}

	for i := n - 1; i > 0; i-- {
		input[0], input[i] = input[i], input[0]
		heapify(input, i, 0)
	}

	return input
}

func heapify(input []int, length int, node int) {
	largest := node
	left := 2*node + 1
	right := 2*node + 2

	if left < length && input[left] > input[largest] {
		largest = left
	}

	if right < length && input[right] > input[largest] {
		largest = right
	}

	if largest != node {
		input[node], input[largest] = input[largest], input[node]
		heapify(input, length, largest)
	}
}
