use std::{fs, str::FromStr, collections::HashMap};

#[test]
fn p1_test() {
    assert_eq!(0, p1(&fs::read_to_string("d6.txt").unwrap()))
}

#[test]
fn p1_ex() {
    let input = "123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i";
    assert_eq!(HashMap::from_iter([
        ("d".to_string(), 72),
        ("e".to_string(), 507),
        ("f".to_string(), 492),
        ("g".to_string(), 114),
        ("h".to_string(), 65412),
        ("i".to_string(), 65079),
        ("x".to_string(), 123),
        ("y".to_string(), 456),
    ]), calc(input))
}

#[test]
fn p2_test() {
    assert_eq!(0, p2(&fs::read_to_string("d6.txt").unwrap()))
}

fn p1(s: &str) -> i32 {
    *calc(s).get("a".to_string()).unwrap()
}

fn p2(s: &str) -> i32 {
    todo!()
}

fn calc(s: &str) -> HashMap<String, i32> {
    todo!()
}