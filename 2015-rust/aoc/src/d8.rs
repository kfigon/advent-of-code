use std::fs;

#[test]
fn p1_example() {
    let d = r#"""
"abc"
"aaa\"aaa"
"\x27""#;
    assert_eq!(12, p1(d));
}

#[test]
fn p2_example() {
    let d = r#"""
"abc"
"aaa\"aaa"
"\x27""#;
    assert_eq!(19, p2(d));
}

#[test]
fn p1_lines() {
    assert_eq!(7, parse_line(r#""aaa\"aaa""#).len());
    assert_eq!(0, parse_line(r#""""#).len());
    assert_eq!(1, parse_line(r#""\x27""#).len());
    assert_eq!(3, parse_line(r#""abc""#).len());
    assert_eq!(2, parse_line(r#""\\"a"#).len());
}

#[test]
fn p1_test() {
    let d = fs::read_to_string("d8.txt").unwrap();
    assert_eq!(1350, p1(&d));
}

#[test]
fn p2_test () {
    let d = fs::read_to_string("d8.txt").unwrap();
    assert_eq!(2085, p2(&d));
}

fn parse_line(s: &str) -> Vec<char> {
    let mut chars = s.chars().collect::<Vec<_>>();
    let mut chars = chars[1..chars.len()-1].iter();

    let mut out = vec![];
    while let Some(&c) = chars.next() {
        if c == '\\' {
            let next = chars.next().unwrap();
            if *next == '"' {
                out.push('"');
            } else if *next == '\\' {
                out.push('\\');
            } else if *next == 'x' {
                match (chars.next(), chars.next()) {
                    (Some(a), Some(b)) => {
                        let v = (*a as u32) * 16 + *b as u32; // todo: wrong conversion
                        out.push(char::from_u32(v).unwrap());
                    }
                    _ => continue,
                }
            }
        } else {
            out.push(c);
        }
    }
    out
}

fn p1(s: &str) -> usize {
    s.lines()
        .map(|v| v.len() - parse_line(v).len())
        .sum()
}

fn p2(s: &str) -> usize {
    s.lines()
        .map(|v| encode(v).len() - v.len())
        .sum()
}

fn encode(s: &str) -> Vec<char> {
    let mut out = vec![];
    out.push('"');
    for c in s.chars() {
        match c {
            '\\' => out.append(&mut vec!['\\','\\']),
            '"' => out.append(&mut vec!['\\','"']),
            a => out.push(a),
        }
    }
    out.push('"');
    out
}