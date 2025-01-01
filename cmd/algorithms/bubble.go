package algorithms

func BubbleSort(input []int) []int {
	for i := len(input); i > 1; i-- {
		for j := 0; j < i-1; j++ {
			if input[j] > input[j+1] {
				input[j], input[j+1] = input[j+1], input[j]
			}
		}
	}
	return input
}
