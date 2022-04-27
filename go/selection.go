package sort

import "golang.org/x/exp/constraints"

func Selection[T constraints.Ordered](arr []T) {
	for step := 0; step < len(arr); step++ {
		minIdx := step

		for i := step + 1; i < len(arr); i++ {
			if arr[i] < arr[minIdx] {
				minIdx = i
			}
		}

		arr[step], arr[minIdx] = arr[minIdx], arr[step]
	}
}
