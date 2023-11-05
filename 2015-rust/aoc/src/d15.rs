use std::{str::FromStr, fs};

const EXAMPLE: &str = "Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3";

#[test]
fn p1_ex() {
    assert_eq!(62842880, p1(EXAMPLE));
}

#[test]
fn p1_test() {
    assert_eq!(18965440, p1(&fs::read_to_string("d15.txt").unwrap()));
}

#[test]
fn p2_ex() {
    assert_eq!(57600000, p2(EXAMPLE));
}

#[test]
fn p2_test() {
    assert_eq!(15862900, p2(&fs::read_to_string("d15.txt").unwrap()));
}

#[derive(Debug, PartialEq)]
struct Ingredient {
    cap: i64,
    dur: i64,
    flav: i64,
    text: i64,
    cal: i64,
}

impl FromStr for Ingredient {
    type Err = &'static str;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let fields = s.split_whitespace()
            .map(|v| v.replace(",", ""))
            .collect::<Vec<_>>();

        match (
            fields.get(2).and_then(|v| v.parse::<i64>().ok()),
            fields.get(4).and_then(|v| v.parse::<i64>().ok()),
            fields.get(6).and_then(|v| v.parse::<i64>().ok()),
            fields.get(8).and_then(|v| v.parse::<i64>().ok()),
            fields.get(10).and_then(|v| v.parse::<i64>().ok()),
        ) {
            (Some(cap), Some(dur), Some(flav), Some(text), Some(cal)) => Ok(Self { cap, dur, flav, text, cal }),
            _ => Err("invalid ingredients")
        }
    }
}

fn parse(s: &str) -> Result<Vec<Ingredient>, &'static str> {
    s.lines().map(|v| v.parse()).collect()
}

// make it lazy iterator to not store everythin in mem
// impl Iterator<Item = Vec<usize>>
pub fn combinations(num_of_ingredients: usize, limit: usize) -> Vec<Vec<usize>> {

    if num_of_ingredients == 4 {
        combinations4(limit)
    } else if num_of_ingredients == 2 {
        let mut out = vec![];
        for i in 0..=limit {
            if i >= limit {
                out.push(vec![i, 0]);
            } else {
                out.push(vec![i, limit-i]);
            }
        }
        out
    } else {
        todo!("make it recursive and accept any number of args")
    }
}
// todo: make this generic and optimised
pub fn combinations4(limit: usize) -> Vec<Vec<usize>> {
    let mut out = vec![];
    for i in 0..=limit {
        for j in 0..=limit {
            for k in 0..=limit {
                for m in 0..=limit {
                    if k+j+i+m == limit {
                        out.push(vec![i,j,k,m]);
                    }
                }
            }
        }
    }
    out
}

fn p1(s: &str) -> i64 {
    calc(s, false)
}

fn p2(s: &str) -> i64 {
    calc(s, true)
}

fn calc(s: &str, incl_calorie_filter: bool) -> i64 {
    let ings = parse(s).unwrap();
    let combs = combinations(ings.len(), 100);

    let mut best = 0i64;

    for c in combs {
        let mut results = [0i64;5];
        
        for (i, &el) in c.iter().enumerate() {
            results[0] += ings[i].cap * el as i64;
            results[1] += ings[i].dur * el as i64;
            results[2] += ings[i].flav * el as i64;
            results[3] += ings[i].text * el as i64;
            results[4] += ings[i].cal * el as i64;
        }
        let current = if !incl_calorie_filter || (incl_calorie_filter && results[4] == 500) { 
            results.into_iter().take(4).map(|v| std::cmp::max(0, v)).product()
        } else {
            0
        };
        best = std::cmp::max(best, current);
    }
    best
}