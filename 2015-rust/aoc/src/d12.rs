use std::fs;

use serde_json::{Value, Map};

#[test]
fn p1_ex_1() {
    let json1 = r#"[1,2,3]"#;
    let json2 = r#"{"a":2,"b":4}"#;
    assert_eq!(p1(json1), 6);
    assert_eq!(p1(json2), 6);
}

#[test]
fn p1_ex_2() {
    let json1 = r#"[[[3]]]"#;
    let json2 = r#"{"a":{"b":4},"c":-1}"#;
    assert_eq!(p1(json1), 3);
    assert_eq!(p1(json2), 3);
}

#[test]
fn p1_ex_3() {
    let json1 = r#"{"a":[-1,1]}"#;
    let json2 = r#"[-1,{"a":1}]"#;
    assert_eq!(p1(json1), 0);
    assert_eq!(p1(json2), 0);
}

#[test]
fn p1_ex_4() {
    let json1 = r#"[]"#;
    let json2 = r#"{}"#;
    assert_eq!(p1(json1), 0);
    assert_eq!(p1(json2), 0);
}

#[test]
fn p2_ex() {
    assert_eq!(p2(r#"[1,2,3]"#), 6);
    assert_eq!(p2(r#"[1,{"c":"red","b":2},3]"#), 4);
    assert_eq!(p2(r#"{"d":"red","e":[1,2,3,4],"f":5}"#), 0);
    assert_eq!(p2(r#"[1,"red",5]"#), 6);
}

#[test]
fn p1_test() {
    assert_eq!(156366, p1(&fs::read_to_string("d12.txt").unwrap()));
}

#[test]
fn p2_test() {
    assert_eq!(96852, p2(&fs::read_to_string("d12.txt").unwrap()));
}

fn p1(s: &str) -> i32 {
    let jval: Value = serde_json::from_str(s).unwrap();
    sum1(jval)
}

fn p2(s: &str) -> i32 {
    let jval: Value = serde_json::from_str(s).unwrap();
    sum2(jval)
}

fn sum1(v: Value) -> i32 {
    match v {
        Value::Null => 0,
        Value::Bool(_) => 0,
        Value::Number(n) => n.as_i64().unwrap() as i32,
        Value::String(_) => 0,
        Value::Array(a) => a.into_iter().map(sum1).sum(),
        Value::Object(o) => o.into_iter().map(|v| sum1(v.1)).sum(),
    }
}

fn sum2(v: Value) -> i32 {
    match v {
        Value::Null => 0,
        Value::Bool(_) => 0,
        Value::Number(n) => n.as_i64().unwrap() as i32,
        Value::String(_) => 0,
        Value::Array(a) => a.into_iter().map(sum2).sum(),
        Value::Object(o) => {
            let mut sum = 0;
            for el in o {
                match &el.1 {
                    Value::String(s) => if *s == "red" {
                        return 0;
                    },
                    _ => sum += sum2(el.1),
                }
            }
            sum
        },
    }
}