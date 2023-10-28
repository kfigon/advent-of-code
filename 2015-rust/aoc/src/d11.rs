use std::collections::{HashSet, HashMap};

#[test]
fn rules_test() {
    assert!(!valid("hijklmmn"));
    assert!(!valid("abbceffg"));
    assert!(!valid("abbcegjk"));
    assert!(!valid("ghijklmn"));
    
    assert!(!valid("abcdefgh"));
    assert!(valid("abcdffaa"));
    assert!(valid("ghjaabcc"));
}

#[test]
fn p1_test() {
    todo!()
}

// Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
// Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
// Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.

// hijklmmn meets the first requirement (because it contains the straight hij) but fails the second requirement requirement (because it contains i and l).
// abbceffg meets the third requirement (because it repeats bb and ff) but fails the first requirement.
// abbcegjk fails the third requirement, because it only has one double letter (bb).

fn valid(s: &str) -> bool {
    let chars = s.chars().collect::<Vec<_>>();
    let invalid_letter = |c: char| -> bool { c.is_uppercase() || !c.is_ascii_alphabetic() || c == 'i' || c == 'o' || c == 'l' };
    let has_pairs = || -> bool {
        let mut substrs: Option<&str> = None;
        for i in 0..chars.len()-1 {
            let x = &s[i..i+2];
            if x.as_bytes()[0] != x.as_bytes()[1] {
                continue;
            }
    
            if substrs.is_none() {
                substrs = Some(x);
            } else if let Some(other) = substrs {
                if other != x {
                    return true;
                }
            }
        }
        false
    };

    chars.len() == 8 &&
    chars.iter().all(|c| !invalid_letter(*c)) &&
    chars.windows(3).any(|c| increasing(c)) &&
    has_pairs()
}

fn increasing(v: &[char]) -> bool {
    next_char(v[0] as u8) == v[1] as u8 && 
    next_char(v[1] as u8) == v[2] as u8 && 
    v[2] as u8 <= b'z'
}

fn next_char(v: u8) -> u8 {
    if v >= b'z' {
        b'a'
    } else {
        v+1
    }
}