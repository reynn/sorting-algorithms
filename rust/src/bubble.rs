use std::cmp::Ord;

pub fn bubble_sort<T: Ord + Copy>(arr: &mut [T]) {
    for i in 0..arr.len() {
        for j in 0..arr.len() - i - 1 {
            if arr[j] > arr[j + 1] {
                arr.swap(j, j + 1);
            }
        }
    }
}

mod test {
    use test_case::test_case;

    #[test_case(&[7, 3, 5, 15, 32, 163, 98], &[3, 5, 7, 15, 32, 98, 164] ; "")]
    fn test_bubble_sort_ints(input: &[i32], expected: &[i32]) {
        let mut arr = input.clone();
        super::bubble_sort(&mut arr);

        assert_eq!(arr, expected);
    }
}
