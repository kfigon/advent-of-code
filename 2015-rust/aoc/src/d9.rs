use std::{collections::{HashMap, HashSet}, fs};

const example: &'static str ="London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141";

#[test]
fn parse_test() {
    let g = parse(example).unwrap();
    assert_eq!(g.0.get("London").unwrap(), &HashMap::from_iter([
        ("Dublin".to_string(), 464),
        ("Belfast".to_string(), 518),
    ]));
}

#[test]
fn p1_ex() {
    assert_eq!(605, p1(example));
}

#[test]
fn p1_test() {
    assert_eq!(141, p1(&fs::read_to_string("d9.txt").unwrap()));
}

#[test]
fn p2_test() {
    assert_eq!(736, p2(&fs::read_to_string("d9.txt").unwrap()));
}

struct Graph(HashMap<String, HashMap<String, usize>>);

fn parse(s: &str) -> Result<Graph, String> {
    let mut out = HashMap::new();
    for line in s.lines() {
        let parts = line.split(" = ").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(format!("invalid line: {line}"));
        }
        let distance = parts[1].parse::<usize>().map_err(|e| format!("invalid distance: {line}: {e}"))?;
        match parts[0].split(" to ").collect::<Vec<_>>()[..] {
            [a,b] => {
                let e = out.entry(a.to_string()).or_insert(HashMap::new());
                e.insert(b.to_string(), distance);
                
                let e2 = out.entry(b.to_string()).or_insert(HashMap::new());
                e2.insert(a.to_string(), distance);
            }
            _ => return Err(format!("invalid destinations: {line}")),
        }
    }

    Ok(Graph(out))
}

fn permutate(vals: &[&str]) -> Vec<Vec<String>> {
    let mut out = vec![];
    if vals.len() == 0 {
        return out;
    }

    for (i, &current) in vals.iter().enumerate() {
        let rest = if i == 0 {
            vals[1..].to_vec()
        } else{
            let mut out = vals[..i].to_vec();
            out.append(&mut vals[i+1..].to_vec());
            out
        };
        
        let mut result = permutate(&rest);
        if result.is_empty() {
            out.push(vec![current.to_string()]);
        } else {
            result.iter_mut().for_each(|c| c.push(current.to_string()));
            out.append(&mut result);
        }
    }

    out
}

fn p1(s: &str) -> usize {
    roads(s).min().unwrap()
}

fn p2(s: &str) -> usize {
    roads(s).max().unwrap()
}

fn roads(s: &str) -> impl Iterator<Item = usize> {
    let g = parse(s).unwrap();

    let cities = g.0.keys().map(|v| v.as_str()).collect::<HashSet<_>>();
    let cities = cities.into_iter().collect::<Vec<_>>();
    let permutations = permutate(&cities);
    let traverse = move |v: Vec<String>| -> usize {
        let mut out = 0;

        for cities in v.windows(2) {
            out += g.0.get(&cities[0]).unwrap().get(&cities[1]).unwrap();
        }

        out
    };
    permutations.into_iter().map(traverse)
}