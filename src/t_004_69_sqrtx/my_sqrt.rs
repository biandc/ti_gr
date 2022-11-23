struct Solution {}

impl Solution {
    pub fn my_sqrt(x: i32) -> i32 {
        0
    }
}


#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_my_sqrt() {
        let x = 0;

        let expected = 0;

        let actual = Solution::my_sqrt(x);

        assert_eq!(expected, actual);
    }
}