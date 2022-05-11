pub fn insertion_sort<T: Ord + Copy>(arr: &mut [T]) {
    for i in 1..arr.len() {
        let key = arr[i];
        let mut j = i - 1;

        while j > 0 && key < arr[j] {
            arr.swap(j + 1, j);
            j = j - 1;
        }
        arr[j + 1] = key;
    }
}

#[test]
fn test_insertion_sort_ints() {
    let mut arr = vec![1, 8, 33, 5, 14, 11];

    insertion_sort(&mut arr);

    assert_eq!(arr, &[1, 5, 8, 11, 14, 33]);
}
