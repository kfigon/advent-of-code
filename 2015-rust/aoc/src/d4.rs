
#[test]
fn p1_test() {
    assert_eq!(254575, p1("bgvyzdsv"));
}

#[test]
fn p1_ex() {
    assert_eq!(609043, p1("abcdef"));
    assert_eq!(1048970, p1("pqrstuv"));
}

#[test]
fn p2_test() {
    assert_eq!(1038736, p2("bgvyzdsv"));
}

pub fn calc_start_with(s: &str, num: usize) -> u64 {
    todo!()
}

fn p1(s: &str) -> u64 {
    calc_start_with(s, 5)
}

fn p2(s: &str) -> u64 {
    calc_start_with(s, 6)
}