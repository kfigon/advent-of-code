use std::{collections::HashMap, vec, fs};

use once_cell::sync::Lazy;
use regex::Regex;

const EXAMPLE: &'static str = "Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.";

#[derive(Debug, PartialEq)]
struct Entry<'a> {
    who: &'a str,
    change: i32,
    target: &'a str
}

impl<'a> TryFrom<&'a str> for Entry<'a> {
    type Error = &'static str;

    fn try_from(s: &'a str) -> Result<Self, Self::Error> {
        static RE: Lazy<Regex> = Lazy::new(|| Regex::new(r"(\w+) would (\w+) (\d+) .* to (\w+)").unwrap());
        
        let res = RE.captures(s).ok_or("invalid input, parts not found")?;

        match (
            res.get(1).map(|v| v.as_str()), 
            res.get(2).map(|v| v.as_str()), 
            res.get(3).map(|v| v.as_str()).map(|v| v.parse::<i32>()), 
            res.get(4).map(|v| v.as_str())) {
                (Some(who), Some("lose"), Some(Ok(change)), Some(target)) => Ok(Entry { who: who, change: -change, target: target }),
                (Some(who), Some("gain"), Some(Ok(change)), Some(target)) => Ok(Entry { who: who, change: change, target: target }),
                (Some(_), Some(_), Some(Err(_)), Some(_)) => Err("invalid number found"),
                (Some(_), Some(_unknown), Some(_), Some(_)) => Err("invalid action, expect 'lose' or 'gain'"),
                _ => Err("invalid input, elements not found"),
        }
    }
}

fn parse(s: &str) -> Graph {
    let es = s.lines()
        .map(|l| l.try_into())
        .collect::<Result<Vec<Entry>, _>>()
        .unwrap();
    Graph::new(es)
}

#[test]
fn parse_test() {
    let e: Result<Entry, _> = "Alice would gain 54 happiness units by sitting next to Bob.".try_into();
    assert_eq!(e, Ok(Entry{who: "Alice", change: 54, target: "Bob"}));

    let e: Result<Entry, _> = "Alice would lose 79 happiness units by sitting next to Carol.".try_into();
    assert_eq!(e, Ok(Entry{who: "Alice", change: -79, target: "Carol"}));

    let e: Result<Entry, _> = "Alice would foo 79 happiness units by sitting next to Carol.".try_into();
    assert_eq!(e, Err("invalid action, expect 'lose' or 'gain'"));

    let e: Result<Entry, _> = "Alice would foo 12a happiness units by sitting next to Carol.".try_into();
    assert_eq!(e, Err("invalid input, parts not found"));

    let e: Result<Entry, _> = "would foo 12 happiness units by sitting next to .".try_into();
    assert_eq!(e, Err("invalid input, parts not found"));
}

#[test]
fn p1_ex() {
    let g = parse(EXAMPLE);
    assert_eq!(330, p1(g));
}

#[test]
fn p1_test() {
    let g = parse(&fs::read_to_string("d13.txt").unwrap());
    assert_eq!(709, p1(g));
}

struct Graph(HashMap<String, HashMap<String, i32>>);

impl Graph {
    fn new(vs: Vec<Entry>) -> Self {
        let mut g: HashMap<String, HashMap<String, i32>> = HashMap::new();
        
        for e in vs {
            let node = g.entry(e.who.to_string()).or_default();
            node.insert(e.target.to_string(), e.change);
        }

        Self(g)
    }

    fn participants(&self) -> Vec<&str> {
        self.0.keys().map(String::as_str).collect()
    }
}

fn permutate(names: &[&str]) -> Vec<Vec<String>> {
    let mut out = vec![];
    if names.len() == 0 {
        return out;
    }

    for (i, &v) in names.iter().enumerate() {
        let rest = if i == 0 {
            names[1..].to_vec()
        } else {
            let mut left = names[..i].to_vec();
            let mut right = names[i+1..].to_vec();
            left.append(&mut right);
            left
        };

        let mut result = permutate(&rest);
        if result.len() == 0 {
            out.push(vec![v.to_string()]);
        } else {
            result.iter_mut().for_each(|l| l.push(v.to_string()));
            out.append(&mut result);
        }
    }

    out
}

fn p1(g: Graph) -> i32 {
    let mut permutations = permutate(&g.participants());
    permutations.iter_mut().for_each(|l| l.push(l[0].to_string()));
    
    let mut sum = 0;
    for perm in permutations {
        let mut subsum = 0;
        for person in perm.windows(2) {
            let x = g.0.get(&person[0]).unwrap().get(&person[1]).unwrap();
            let y = g.0.get(&person[1]).unwrap().get(&person[0]).unwrap();

            subsum = subsum + *y + *x;
        }

        if subsum > sum {
            sum = subsum;
        }
    }
    sum
}