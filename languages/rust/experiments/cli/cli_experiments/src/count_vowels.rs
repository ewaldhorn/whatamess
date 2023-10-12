fn get_count(string: &str) -> usize {
    let vowels = ['a', 'e', 'i', 'o', 'u'];
    string.to_ascii_lowercase().chars().filter(|character| vowels.iter().any(|vowel|vowel==character)).count()
}

#[test]
fn test_get_count() {
    let input = "This contains eight vowels.";
    assert_eq!(get_count(input), 8)
}
