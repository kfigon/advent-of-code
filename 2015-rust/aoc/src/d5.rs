use std::{fs, collections::HashSet};

#[test]
fn p1_test() {
    assert_eq!(258, p1(&fs::read_to_string("d5.txt").unwrap()))
}

#[test]
fn p2_test() {
    assert_eq!(53, p2(&fs::read_to_string("d5.txt").unwrap()))
}

#[test]
fn is_nice_p2_test() {
    assert!(!is_nice_p2("xyxaaac"));
    assert!(is_nice_p2("xyxaaaa"));
    assert!(is_nice_p2("qjhvhtzxzqqjkmpb"));
    assert!(is_nice_p2("xxyxx"));

    assert!(!is_nice_p2("aaa"));
    assert!(!is_nice_p2("uurcxstgmygtbstg"));
    assert!(!is_nice_p2("ieodomkazucvgmuy"));
}

fn is_nice_p1(s: &str) -> bool {
    let chars = s.chars().collect::<Vec<_>>();
    let vowels = chars.iter()
        .filter(|&&c| c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u')
        .count();

    let pairs = chars.windows(2).any(|c| c[0] == c[1]);
    let forbidden = chars.windows(2).any(|c| (c[0] == 'a' && c[1] == 'b') || (c[0] == 'c' && c[1] == 'd') || (c[0] == 'p' && c[1] == 'q') || (c[0] == 'x' && c[1] == 'y'));
    
    return vowels >= 3 && pairs && !forbidden;
}

fn is_nice_p2(s: &str) -> bool {
    let chars = s.chars().collect::<Vec<_>>();
    let repeats_with_letter_between = chars.windows(3).any(|c| c[0] == c[2]);

    let mut pair_twice = false;
    for i in 0..chars.len()-1 {
        for j in i+2..chars.len()-1 {
            if chars[i] == chars[j] && chars[i+1] == chars[j+1] {
                pair_twice = true;
                break;
            }
        }
    }
    
    repeats_with_letter_between && pair_twice
}

fn p1(s: &str) -> usize {
    s.lines().filter(|line| is_nice_p1(line)).count()
}

fn p2(s: &str) -> usize {
    s.lines().filter(|line| is_nice_p2(line)).count()
}