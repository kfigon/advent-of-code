use std::{collections::HashMap, fmt::format};

#[test]
fn parse_test() {
    let d = "London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141";

    let g = parse(d).unwrap();
    assert_eq!(g.0.get("London").unwrap(), &HashMap::from_iter([
        ("Dublin".to_string(), 464),
        ("Belfast".to_string(), 518),
    ]));
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