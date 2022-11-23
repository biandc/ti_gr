struct Solution {}

impl Solution {
    pub fn search_range(nums: Vec<i32>, target: i32) -> Vec<i32> {
        vec![0, 0]
    }
}


#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_search_range() {
        let nums = vec![5, 7, 7, 8, 8, 10];
        let target = 0;

        let expected = vec![0, 0];

        let actual = Solution::search_range(nums, target);

        assert_eq!(expected, actual);
    }
}