use std::{collections::HashMap, fs};

fn parse_line(s: &str) -> Option<(i32, HashMap<&str, i32>)> {
    let parts = s.split_once(": ")?;

    let id = parts.0.split(" ").collect::<Vec<_>>();
    let id = id.get(1).and_then(|v| v.parse::<i32>().ok())?;

    let props = parts.1.split(", ").collect::<Vec<_>>();
    let d = props.iter()
        .map(|&v| v.split(": ").collect::<Vec<_>>())
        .map(|pair| match (pair.get(0), pair.get(1).and_then(|v| v.parse::<i32>().ok())) {
            (Some(&key), Some(num)) => Some((key, num)),
            _ => None,
        })
        .collect::<Option<Vec<_>>>()?;
    
    Some((id, HashMap::from_iter(d)))
}

#[test]
fn parse_test() {
    let data = "Sue 4: goldfish: 10, akitas: 2, perfumes: 9";
    assert_eq!(parse_line(data).unwrap(), (4, HashMap::from_iter([
        ("goldfish", 10),
        ("akitas", 2),
        ("perfumes", 9),
    ])))
}

fn input() -> HashMap<&'static str, i32> {
    HashMap::from_iter([
        ("children", 3),
        ("cats", 7),
        ("samoyeds", 2),
        ("pomeranians", 3),
        ("akitas", 0),
        ("vizslas", 0),
        ("goldfish", 5),
        ("trees", 3),
        ("cars", 2),
        ("perfumes", 1),
    ])
}

#[test]
fn p1_test() {
    assert_eq!(40, p1(&fs::read_to_string("d16.txt").unwrap()));
}

#[test]
fn p2_test() {
    assert_eq!(241, p2(&fs::read_to_string("d16.txt").unwrap()));
}

fn p1(s: &str) -> i32 {
    let is_valid_candidate = |d: &HashMap<&str, i32>, model: &HashMap<&str, i32>| -> bool {
        model.iter()
            .all(|(&key, &expected_val)| match d.get(key) {
                None => true,
                Some(&v) => v == expected_val,
            })
    };

    solve(s, is_valid_candidate)
}

fn p2(s: &str) -> i32 {
    let is_valid_candidate = |d: &HashMap<&str, i32>, model: &HashMap<&str, i32>| -> bool {
        model.iter()
            .all(|(&key, &expected_val)| match d.get(key) {
                None => true,
                Some(&v) => match key {
                    "cats" | "trees" => expected_val < v,
                    "pomeranians" | "goldfish" =>  expected_val > v,
                    _ => v == expected_val,
                }
            })
    };

    solve(s, is_valid_candidate)
}

fn solve(s: &str, candidate_fn: impl Fn(&HashMap<&str, i32>, &HashMap<&str, i32>) -> bool) -> i32 {
    let data = s.lines()
        .map(parse_line)
        .collect::<Option<Vec<_>>>()
        .unwrap();
    
    let model = input();

    let candidate = data.iter()
        .find(|(_, d)| candidate_fn(d, &model));

    candidate.unwrap().0
}