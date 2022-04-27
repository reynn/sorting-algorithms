package sort

import "golang.org/x/exp/constraints"

func Insertion[T constraints.Ordered](arr []T) {
	for i := 0; i < len(arr); i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && key < arr[j] {
			arr[j+1] = arr[j]
			j = j - 1
		}

		arr[j+1] = key
	}
}
