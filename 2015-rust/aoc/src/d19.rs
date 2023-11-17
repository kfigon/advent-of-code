use std::{collections::HashMap, fs};

const EXAMPLE1: &'static str = "H => HO
H => OH
O => HH

HOH";

const EXAMPLE2: &'static str = "H => HO
H => OH
O => HH

HOHOHO";

#[test]
fn parse_ex() {
    assert_eq!(parse(EXAMPLE1), Ok((
        HashMap::from_iter([
            ("H", "HO"),
            ("H", "OH"),
            ("O", "HH"),
        ]),
        "HOH")
    ))
}

#[test]
fn p1_ex_1() {
    assert_eq!(4, p1(EXAMPLE1));
}

#[test]
fn p1_ex_2() {
    assert_eq!(7, p1(EXAMPLE2));
}

#[test]
fn p1_test() {
    todo!();
    assert_eq!(1, p1(&fs::read_to_string("d19.txt").unwrap()));
}

fn parse(s: &str) -> Result<(HashMap<&str, &str>, &str), &'static str> {
    let parts = s.split_once("\n\n");
    let (mappings, formula) = match parts {
        Some(v) => (v.0,v.1),
        None => return Err("parts not found"),
    };

    let mappings = mappings.lines()
        .map(|line| line.split_once(" => "))
        .collect::<Option<HashMap<&str,&str>>>().ok_or("invalid mapping found")?;

    Ok((mappings, formula))
}

fn p1(s: &str) -> usize {
    todo!()
}