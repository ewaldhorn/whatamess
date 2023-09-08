fn main() {
    let mut a = format!("");
    std::io::stdin().read_line(&mut a).unwrap();
    let mut b = a.trim_matches('\n').to_string();
    
    a = format!("");
    for c in b.into_bytes() {
        a += &format!("{:b}", c);
    }
    b = format!("");
    let mut d = a.chars().take(1).last().unwrap();
    if d == '0' {
        b += "00 0";
    } else {
        b += "0 0";
    }
    for c in a.chars().skip(1) {
        if c == d {
            b += "0";
        } else {
            d = c;
            b += " ";
            if d == '0' {
                b += "00 0";
            } else {
                b += "0 0";
            }
        }
    }
    println!("{b}")
}

