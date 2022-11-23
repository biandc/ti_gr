struct Solution {}

impl Solution {
    pub fn search_insert(nums: Vec<i32>, target: i32) -> i32 {
        1
    }
}


#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_search_insert() {
        let nums = vec![-1, 3, 5, 7];
        let target = 0;

        let expected = 1;

        let actual = Solution::search_insert(nums, target);

        assert_eq!(expected, actual);
    }
}
