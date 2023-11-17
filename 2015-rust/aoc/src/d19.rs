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
fn p1_ex_3() {
    // // SiOSi
    // mHSi
    // SiHm
    assert_eq!(3, p1(EXAMPLE3));
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
        .collect::<Option<HashMap<&str,&str>>>()
        .ok_or("invalid mapping found")?;

    Ok((mappings, formula))
}

fn p1(s: &str) -> usize {
    let (mappings, formula) = parse(s).unwrap();
    let mut unique: HashSet<String> = HashSet::new();

    for (&k,&v) in &mappings {        
        let mut i = 0;
        while i < formula.len() {
            i = match substr(formula, k, i) {
                Ok(found_idx) => {
                    let mut molecule = String::new();
                    molecule.push_str(&formula[..i]);
                    molecule.push_str(v);
                    molecule.push_str(&formula[i+k.len()..]);


                    unique.insert(molecule);
                    found_idx+1
                },
                Err(next_idx) => next_idx,
            };
        }
    }

    unique.len()
}

fn substr(s: &str, sub: &str, start_idx: usize) -> Result<usize, usize> {
    let mut i = start_idx;
    while i < s.len() {
        match s.get(i..i+sub.len()) {
            Some(v) if v == sub => return Ok(i),
            _ => i+=1,
        };
    }
    Err(i)
}