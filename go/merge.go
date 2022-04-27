package sort

import (
	"golang.org/x/exp/constraints"
)

func Merge[T constraints.Ordered](arr []T) []T {
	num := len(arr)

	if num == 1 {
		return arr
	}

	mid := int(num / 2)
	var (
		left  = make([]T, mid)
		right = make([]T, num-mid)
	)
	for i := 0; i < num; i++ {
		if i < mid {
			left[i] = arr[i]
		} else {
			right[i-mid] = arr[i]
		}
	}

	return m(Merge(left), Merge(right))
}

func m[T constraints.Ordered](left, right []T) []T {
	result := make([]T, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return result
}
