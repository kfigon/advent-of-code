use std::{str::FromStr, collections::HashMap};

const EXAMPLE: &str = "Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3";

#[test]
fn p1_ex() {
    todo!()
}

#[derive(Debug, PartialEq)]
struct Ingredient {
    cap: i32,
    dur: i32,
    flav: i32,
    text: i32,
    cal: i32,
}

impl FromStr for Ingredient {
    type Err = &'static str;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let fields = s.split_whitespace()
            .map(|v| v.replace(",", ""))
            .collect::<Vec<_>>();

        match (
            fields.get(2).and_then(|v| v.parse::<i32>().ok()),
            fields.get(4).and_then(|v| v.parse::<i32>().ok()),
            fields.get(6).and_then(|v| v.parse::<i32>().ok()),
            fields.get(8).and_then(|v| v.parse::<i32>().ok()),
            fields.get(10).and_then(|v| v.parse::<i32>().ok()),
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
    todo!()
}

fn p1(s: &str) -> i64 {
    let ings = parse(s).unwrap();
    let combs = combinations(ings.len(), 100);

    let mut mapping: HashMap<usize, Box<dyn Fn(&Ingredient) -> i32>> = HashMap::new();
    mapping.insert(0, Box::new(|i: &Ingredient| -> i32 { i.cap }));
    mapping.insert(1, Box::new(|i: &Ingredient| -> i32 { i.dur }));
    mapping.insert(2, Box::new(|i: &Ingredient| -> i32 { i.flav }));
    mapping.insert(3, Box::new(|i: &Ingredient| -> i32 { i.text }));
    let mut best = 0;

    for c in combs {
        for (i, &portion) in c.iter().enumerate() {
            let fun = mapping.get(&i).unwrap();
            
        }
    }
    best
}