package algorithms

func Merge_sort(input []int) []int {
	if len(input) <= 1 {
		return input
	} else {

		mid := len(input) / 2

		left := Merge_sort(input[:mid])
		right := Merge_sort(input[mid:])
		return merge(left, right)
	}

}

func merge(left []int, right []int) []int {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {

		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for i < len(left) {
		result = append(result, left[i])
		i++
	}
	for j < len(right) {
		result = append(result, right[j])
		j++
	}

	return result
}
