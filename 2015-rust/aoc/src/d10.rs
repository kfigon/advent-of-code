#[test]
fn p1_ex() {
    assert_eq!(next("1"), "11".to_string());
    assert_eq!(next("11"), "21".to_string());
    assert_eq!(next("21"), "1211".to_string());
    assert_eq!(next("1211"), "111221".to_string());
    assert_eq!(next("111221"), "312211".to_string());
}

#[test]
fn p1_test() {
    assert_eq!(solve("3113322113", 40), 329356);
}

#[test]
#[ignore = "long time"]
fn p2_test() {
    assert_eq!(solve("3113322113", 50), 4666278);
}

fn next(s: &str) -> String {
    let chars = s.chars().map(|v| v.to_digit(10)).collect::<Option<Vec<u32>>>();
    let digits = match chars {
        None => return "".to_string(),
        Some(v) => if v.len() == 0 {
            return "".to_string();
        } else {
            v
        },
    };

    let mut last: (u32, i32) = (digits[0], 1);
    let mut out = String::new();
    
    for &c in &digits[1..] {
        if last.0 == c {
            last.1+=1;
        } else {
            out.push_str(&format!("{}{}", last.1, last.0.to_string()));
            last = (c, 1);
        }
    }

    out.push_str(&format!("{}{}", last.1, last.0.to_string()));

    out
}

fn solve(s: &str, len: usize) -> usize {
    let mut out = s.to_string();
    for _ in 0..len {
       out = next(&out);
    }

    out.len()
}