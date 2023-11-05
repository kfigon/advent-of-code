use std::{str::FromStr};

const EXAMPLE: &str = "Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3";

#[test]
fn p1_ex() {
    assert_eq!(62842880, p1(EXAMPLE));
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
fn combinations(num_of_ingredients: usize, limit: usize) -> Vec<Vec<usize>> {
    let mut out = vec![];
    for i in 0..=limit {
        // todo: how to make it flexible?
        out.push(vec![i, limit-i]);
    }
    out
}

fn p1(s: &str) -> i64 {
    let ings = parse(s).unwrap();
    let combs = combinations(ings.len(), 100);

    let mut best = 0i64;

    for c in combs {
        let mut results = [0i64;4];
        
        for (i, &el) in c.iter().enumerate() {
            results[0] += ings[i].cap * el as i64;
            results[1] += ings[i].dur * el as i64;
            results[2] += ings[i].flav * el as i64;
            results[3] += ings[i].text * el as i64;
        }
        let current = results.map(|v| std::cmp::max(0, v)).iter().product();
        best = std::cmp::max(best, current);
    }
    best
}