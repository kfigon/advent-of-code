use std::fs;

#[test]
fn p1_test() {
    assert_eq!(280, p1(&fs::read_to_string("d1.txt").unwrap()))
}

#[test]
fn p2_test() {
    assert_eq!(1797, p2(&fs::read_to_string("d1.txt").unwrap()))
}

fn p1(s: &str) -> i32 {
    s.chars()
        .fold(0, |acc, v| match v {
            ')' => acc-1,
            '(' => acc+1,
            _ => todo!("invalid char {v}"), 
        })
}

fn p2(s: &str) -> usize {
    let mut current = 0;
    for (i, v) in s.char_indices() {
        if v == ')' {
            current -= 1;
        } else if v == '(' {
            current += 1;
        }
        
        if current == -1 {
            return i+1;
        }
    }
    return 0;
}