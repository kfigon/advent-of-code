#[test]
fn rules_test() {
    let adj = |s: &str| -> Vec<char> { s.chars().collect()};
    assert!(!valid(&adj("hijklmmn")));
    assert!(!valid(&adj("abbceffg")));
    assert!(!valid(&adj("abbcegjk")));
    assert!(!valid(&adj("ghijklmn")));
    
    assert!(!valid(&adj("abcdefgh")));
    assert!(valid(&adj("abcdffaa")));
    assert!(valid(&adj("ghjaabcc")));
}

#[test]
fn p1_ex() {
    assert_eq!(p1("abcdefgh"), "abcdffaa".to_string());
    assert_eq!(p1("ghijklmn"), "ghjaabcc".to_string());
}

#[test]
fn p1_test() {
    assert_eq!(p1("vzbxkghb"), "vzbxxyzz".to_string());
}

// Passwords must include one increasing straight of at least three letters, like abc, bcd, cde, and so on, up to xyz. They cannot skip letters; abd doesn't count.
// Passwords may not contain the letters i, o, or l, as these letters can be mistaken for other characters and are therefore confusing.
// Passwords must contain at least two different, non-overlapping pairs of letters, like aa, bb, or zz.

// hijklmmn meets the first requirement (because it contains the straight hij) but fails the second requirement requirement (because it contains i and l).
// abbceffg meets the third requirement (because it repeats bb and ff) but fails the first requirement.
// abbcegjk fails the third requirement, because it only has one double letter (bb).

fn valid(chars: &Vec<char>) -> bool {
    let invalid_letter = |c: char| -> bool { c.is_uppercase() || !c.is_ascii_alphabetic() || c == 'i' || c == 'o' || c == 'l' };
    let has_pairs = || -> bool {
        let mut substrs: Option<(char,char)> = None;
        for i in 0..chars.len()-1 {
            let x = (*chars.get(i).unwrap(), *chars.get(i+1).unwrap());
            if x.0 != x.1 {
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
    
    let increasing = |v: &[char]| -> bool {
        next_char(v[0] as u8) == v[1] as u8 && 
        next_char(v[1] as u8) == v[2] as u8 && 
        v[2] as u8 <= b'z'
    };

    chars.len() == 8 &&
    chars.iter().all(|c| !invalid_letter(*c)) &&
    chars.windows(3).any(|c| increasing(c)) &&
    has_pairs()
}

fn next_char(v: u8) -> u8 {
    if v >= b'z' {
        b'a'
    } else {
        v+1
    }
}

fn next_pass(mut chars: Vec<char>) -> Vec<char> {
    let mut i = chars.len()-1;

    chars[i] = next_char(chars[i] as u8) as char;
    let mut carry = chars[i] == 'a';
    i-=1;

    while carry {
        chars[i] = next_char(chars[i] as u8) as char;
        carry = chars[i] == 'a';
        i -= 1;
        if i == 0 {
            break;
        }
    }

    chars
}

pub fn p1(s: &str) -> String {
    let mut chars = s.chars().collect::<Vec<_>>();
    chars = next_pass(chars);

    while !valid(&chars) {
        chars = next_pass(chars);
    }

    chars.into_iter().fold(String::new(), |mut acc, v| {
        acc.push(v);
        acc
    })
}