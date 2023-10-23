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

enum Op {
    Load(Signal, String),
    And(Signal, Signal, String),
    Or(Signal, Signal, String),
    Lshift(Signal, i16, String),
    Rshift(Signal, i16, String),
    Not(Signal, String),
}
// 123 -> x
// x AND y -> z
// x OR y -> z
// p LSHIFT 2 -> q
// p RSHIFT 2 -> q
// NOT e -> f
impl FromStr for Op {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts = s.split(" -> ").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(format!("signal not routed: {s}"));
        }
        let first_part = parts[0];
        let target = parts[1];

        
    }
}

enum Signal {
    Wire(String),
    Constant(i16)
}

impl FromStr for Signal {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        if s.trim().is_empty() {
            return Err(format!("empty signal: {s}"));
        }
        match s.parse::<i16>() {
            Ok(i) => Ok(Signal::Constant(i)),
            Err(_) => Ok(Signal::Wire(s.to_string())),
        }
    }
}

#[test]
fn p2_test() {
    assert_eq!(0, p2(&fs::read_to_string("d6.txt").unwrap()))
}

fn p1(s: &str) -> u16 {
    *calc(s).get("a").unwrap()
}

fn p2(s: &str) -> u16 {
    todo!()
}

fn calc(s: &str) -> HashMap<String, u16> {
    todo!()
}