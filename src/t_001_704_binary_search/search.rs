struct Solution {}

impl Solution {
    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        // TODO: codes

        4
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_search() {
        let nums = vec![-1, 0, 3, 5, 9, 12, 18];
        let target = 9;

        let expected = 4;

        let actual = Solution::search(nums, target);

        assert_eq!(expected, actual);
    }
}
