struct Solution {}

impl Solution {
    pub fn is_perfect_square(num: i32) -> bool {
        true
    }
}


#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn test_is_perfect_square() {
        let num = 8;

        let expected = false;

        let actual = Solution::is_perfect_square(num);

        assert_eq!(expected, actual);
    }
}