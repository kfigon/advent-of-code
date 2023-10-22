extern crate md5;

#[test]
#[ignore = "long time"]
fn p1_test() {
    assert_eq!(254575, p1("bgvyzdsv"));
}

#[test]
#[ignore = "long time"]
fn p1_ex() {
    assert_eq!(609043, p1("abcdef"));
    assert_eq!(1048970, p1("pqrstuv"));
}

#[test]
#[ignore = "long time"]
fn p2_test() {
    assert_eq!(1038736, p2("bgvyzdsv"));
}

pub fn calc_start_with(s: &str, num: usize) -> Option<u64> {
    for i in 0..u64::MAX {
        let digest = md5::compute(format!("{s}{i}").as_bytes());
        let hex = format!("{:?}", digest);
        if hex.chars().take(num).all(|v| v == '0') {
            return Some(i);
        }
    }
    None
}

fn p1(s: &str) -> u64 {
    calc_start_with(s, 5).unwrap()
}

fn p2(s: &str) -> u64 {
    calc_start_with(s, 6).unwrap()
}