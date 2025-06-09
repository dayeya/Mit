package main

func mergeBoth(left []int, right []int) []int {
	merged := make([]int, 0, len(left)+len(right))
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			merged = append(merged, left[i])
			i++
		} else {
			merged = append(merged, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		merged = append(merged, left[i])
	}
	for ; j < len(right); j++ {
		merged = append(merged, right[j])
	}

	return merged
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return mergeBoth(left, right)
}
