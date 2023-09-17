use unicode_segmentation::UnicodeSegmentation;

pub fn reverse(input: &str) -> String {
    let mut tmp = input.graphemes(true).collect::<Vec<&str>>();
    tmp.reverse();
    tmp.join("")
}
