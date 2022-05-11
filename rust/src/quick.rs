pub fn quicksort<T: Ord + Copy>(arr: &mut [T]) {
    let len = arr.len();
    quick_sort(arr, 0, len - 1);
}

fn quick_sort<T: Ord + Copy>(arr: &mut [T], low: usize, high: usize) {
    if low < high {
        let p = partition(arr, low, high);
        quick_sort(arr, low, p - 1);
        quick_sort(arr, p + 1, high);
    }
}

fn partition<T: Ord + Copy>(arr: &mut [T], low: usize, high: usize) -> usize {
    let pivot = high as usize;
    let mut store_index = if low == 0 { 0 } else { low - 1 };
    let mut last_index = high;

    loop {
        store_index += 1;
        while arr[store_index] < arr[pivot] {
            store_index += 1;
        }
        last_index -= 1;
        while last_index > 0 && arr[last_index] > arr[pivot] {
            last_index -= 1;
        }
        if store_index >= last_index {
            break;
        } else {
            arr.swap(store_index, last_index);
        }
    }
    arr.swap(store_index, pivot);
    store_index
}

#[test]
fn test_quick_sort_ints() {
    let mut arr = vec![1, 8, 33, 5, 14, 11];

    quicksort(&mut arr);

    assert_eq!(arr, &[1, 5, 8, 11, 14, 33]);
}
