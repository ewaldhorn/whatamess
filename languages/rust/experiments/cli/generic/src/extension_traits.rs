pub trait Palindrome {
    fn is_palindrome(&self) -> bool;
}

impl Palindrome for String {
    fn is_palindrome(&self) -> bool {
        let s = self.chars().collect::<Vec<char>>();
        s == s.iter().rev().cloned().collect::<Vec<char>>()
    }
}
