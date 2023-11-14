use std::{num::ParseIntError, fs};

#[test]
fn p1_ex() {
    let limit = 25;
    let example = "20
15
10
5
5";

    assert_eq!(4, p1(example, limit));
}

#[test]
fn p1_test() {
    assert_eq!(654, p1(&fs::read_to_string("d17.txt").unwrap(), 150));
}

#[test]
fn p2_test() {
    assert_eq!(57, p2(&fs::read_to_string("d17.txt").unwrap(), 150));
}

fn parse(s: &str) -> Result<Vec<i32>, ParseIntError> {
    s.lines().map(|v| v.parse::<i32>()).collect()
}

// brute force :(
fn p1(s: &str, limit: i32) -> usize {
    calc_combinations(s, limit).count()
}

fn calc_combinations(s: &str, limit: i32) -> impl Iterator<Item = Vec<i32>> {
    let d = parse(s).unwrap();

    let combs = combinate(&d);

    combs.into_iter()
        .filter(move |v| v.iter().sum::<i32>() == limit)
}

// order does not matter (in permutations it does)
fn combinate(d: &[i32]) -> Vec<Vec<i32>> {
    if d.is_empty() {
        return vec![vec![]];
    }

    let first = *d.first().unwrap();
    let rest = &d[1..];

    let combs_without_first = combinate(rest);
    
    let mut out = vec![];
    
    // take elements without the first and then with
    for c in &combs_without_first {
        out.push(c.clone());
    }

    for mut c in combs_without_first {
        c.push(first);
        out.push(c);
    }

    out
}

fn p2(s: &str, limit: i32) -> usize {
    let valid_combinations = calc_combinations(s, limit).collect::<Vec<Vec<i32>>>();
    let min_len = valid_combinations.iter().map(|v| v.len()).min().unwrap();
    valid_combinations.iter().filter(|v| v.len() == min_len).count()
}