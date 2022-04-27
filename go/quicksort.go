package sort

import "golang.org/x/exp/constraints"

func Quick[T constraints.Ordered](arr []T, low, high int) {
	if low > high {
		return
	}

	part := partition(arr, low, high)

	Quick(arr, low, part-1)
	Quick(arr, part+1, high)
}

func partition[T constraints.Ordered](arr []T, low, high int) int {
	pivot := arr[high]
	i := low - 1

	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			// if element is smaller than the the pivot swap with greater element pointed by i
			i++

			// swapping element at `i` with elemnt at `j`
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[i+1], arr[high] = arr[high], arr[i+1]

	return i + 1
}
