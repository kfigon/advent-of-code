use std::{collections::{HashMap, HashSet}, fs};

const EXAMPLE1: &'static str = "H => HO
H => OH
O => HH

HOH";

const EXAMPLE2: &'static str = "H => HO
H => OH
O => HH

HOHOHO";

const EXAMPLE3: &'static str = "H => O
Si => m

SiHSi";


#[test]
fn parse_ex() {
    assert_eq!(parse(EXAMPLE1), Ok((
        HashMap::from_iter([
            ("H", vec!["HO", "OH"]),
            ("O", vec!["HH"]),
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
fn p1_ex_3() {
    // // SiOSi
    // mHSi
    // SiHm
    assert_eq!(3, p1(EXAMPLE3));
}

#[test]
fn substr_test() {
    assert_eq!(substr("fooBarfoo", "foo", 0), Some(0));
    assert_eq!(substr("fooBarfoo", "foo", 1), Some(6));
    assert_eq!(substr("fooBarfoo", "foo", 2), Some(6));
    assert_eq!(substr("fooBarfoo", "foo", 4), Some(6));
    assert_eq!(substr("fooBarfoo", "Bar", 0), Some(3));
    assert_eq!(substr("fooBarfoo", "Bar", 1), Some(3));
    assert_eq!(substr("fooBarfoo", "Bar", 3), Some(3));
    assert_eq!(substr("fooBarfoo", "Bar", 2), Some(3));
    
    assert_eq!(substr("fooBarfoo", "Bar", 4), None);
    assert_eq!(substr("fooBarfoo", "bar", 0), None);
    assert_eq!(substr("fooBarfoo", "foo", 7), None);
}

#[test]
fn p1_test() {
    assert_eq!(518, p1(&fs::read_to_string("d19.txt").unwrap()));
}

fn parse(s: &str) -> Result<(HashMap<&str, Vec<&str>>, &str), &'static str> {
    let parts = s.split_once("\n\n");
    let (mappings, formula) = match parts {
        Some(v) => (v.0,v.1),
        None => return Err("parts not found"),
    };

    let mappings = mappings.lines()
        .map(|line| line.split_once(" => "))
        .collect::<Option<Vec<(&str,&str)>>>()
        .ok_or("invalid mapping found")?;

    let out = mappings.iter()
        .fold(HashMap::new(), |mut acc, v| {
            let e: &mut Vec<&str> = acc.entry(v.0).or_default();
            e.push(v.1);
            acc
        });
    Ok((out, formula))
}

fn p1(s: &str) -> usize {
    let (mappings, formula) = parse(s).unwrap();

    let unique = mappings.iter()
        .flat_map(|v| v.1.iter().map(move |el| (v.0, el))) // str,Vec<key> -> flat Vec (str, key)
        .fold(HashSet::new(), |mut unique, (&k, &v)| {
            let mut i = 0;
            while i < formula.len() {
                i = match substr(formula, k, i) {
                    Some(found_idx) => {
                        let mut molecule = String::new();
                        molecule.push_str(&formula[..found_idx]);
                        molecule.push_str(v);
                        molecule.push_str(&formula[found_idx+k.len()..]);
                        
                        unique.insert(molecule);
                        found_idx+1
                    },
                    None => break,
                };
            }
            unique
        });

    unique.len()
}

fn substr(s: &str, sub: &str, start_idx: usize) -> Option<usize> {
    let mut i = start_idx;
    while i < s.len() {
        match s.get(i..i+sub.len()) {
            Some(v) if v == sub => return Some(i),
            _ => i+=1,
        };
    }
    None
}