use std::{fs, str::FromStr, collections::{HashMap}};

#[test]
fn p1_test() {
    assert_eq!(46065, p1(&fs::read_to_string("d7.txt").unwrap()))
}

#[test]
fn p2_test() {
    assert_eq!(14134, p2(&fs::read_to_string("d7.txt").unwrap()))
}

#[test]
fn p1_ex() {
    let input = "123 -> x
456 -> y
x AND y -> d
x OR y -> e
x LSHIFT 2 -> f
y RSHIFT 2 -> g
NOT x -> h
NOT y -> i";
    assert_eq!(HashMap::from_iter([
        ("d".to_string(), 72),
        ("e".to_string(), 507),
        ("f".to_string(), 492),
        ("g".to_string(), 114),
        ("h".to_string(), 65412),
        ("i".to_string(), 65079),
        ("x".to_string(), 123),
        ("y".to_string(), 456),
    ]), calc(input))
}

#[derive(Clone)]
enum Cmd {
    Load(Signal),
    And(Signal, Signal),
    Or(Signal, Signal),
    Lshift(Signal, u16),
    Rshift(Signal, u16),
    Not(Signal),
}

#[derive(Clone)]
struct Op(Cmd, String);

// 123 -> x
// x AND y -> z
// x OR y -> z
// p LSHIFT 2 -> q
// p RSHIFT 2 -> q
// NOT e -> f
impl FromStr for Op {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let parts = s.split(" -> ").collect::<Vec<_>>();
        if parts.len() != 2 {
            return Err(format!("signal not routed: {s}"));
        }
        let target = parts[1].to_string();

        let cmd_parts = parts[0].split_whitespace().collect::<Vec<_>>();
        match cmd_parts[..] {
            [v] => Ok(Op(Cmd::Load(v.parse()?), target)),
            [a, "AND", b] => Ok(Op(Cmd::And(a.parse()?, b.parse()?), target)),
            [a, "OR", b] => Ok(Op(Cmd::Or(a.parse()?, b.parse()?), target)),
            [a, "LSHIFT", b] => Ok(Op(Cmd::Lshift(a.parse()?, b.parse().map_err(|_| format!("invalid shift val: {s}"))?), target)),
            [a, "RSHIFT", b] => Ok(Op(Cmd::Rshift(a.parse()?, b.parse().map_err(|_| format!("invalid shift val: {s}"))?), target)),
            ["NOT", b] => Ok(Op(Cmd::Not(b.parse()?), target)),
            _ => Err(format!("invalid cmd {s}")),
        }
    }
}

#[derive(Clone)]
enum Signal {
    Wire(String),
    Constant(u16)
}

impl FromStr for Signal {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        if s.trim().is_empty() {
            return Err(format!("empty signal: {s}"));
        }
        match s.parse::<u16>() {
            Ok(i) => Ok(Signal::Constant(i)),
            Err(_) => Ok(Signal::Wire(s.to_string())),
        }
    }
}

struct Graph(HashMap<String, Cmd>);

impl Graph {
    fn new(ops: Vec<Op>) -> Self {
        let mut g = HashMap::new();
        for o in ops {
            g.insert(o.1, o.0);
        }
        Self(g)
    }

    fn process(&self) -> HashMap<String, u16> {
        let mut out = HashMap::new();

        for (k,o) in &self.0 {
            self.visit(k, o, &mut out);
        }

        out
    }

    fn visit(&self, k: &str, c: &Cmd, visited: &mut HashMap<String, u16>) -> u16 {
        if let Some(v) = visited.get(k) {
            return *v;
        }
        
        let r = match c {
            Cmd::Load(v) => match v {
                Signal::Wire(w) => self.visit(w, self.0.get(w).unwrap(), visited),
                Signal::Constant(c) => *c,
            },
            Cmd::Not(v) => match v {
                Signal::Wire(w) => !self.visit(w, self.0.get(w).unwrap(), visited),
                Signal::Constant(c) => *c,
            },
            Cmd::And(a, b) => match (a,b) {
                (Signal::Wire(w1), Signal::Wire(w2)) => self.visit(w1, self.0.get(w1).unwrap(), visited) & self.visit(w2, self.0.get(w2).unwrap(), visited),
                (Signal::Wire(w), Signal::Constant(c)) => self.visit(w, self.0.get(w).unwrap(), visited) & c,
                (Signal::Constant(c), Signal::Wire(w)) => c & self.visit(w, self.0.get(w).unwrap(), visited),
                (Signal::Constant(c1), Signal::Constant(c2)) => c1 & c2,
            },
            Cmd::Or(a,b) => match (a,b) {
                (Signal::Wire(w1), Signal::Wire(w2)) => self.visit(w1, self.0.get(w1).unwrap(), visited) | self.visit(w2, self.0.get(w2).unwrap(), visited),
                (Signal::Wire(w), Signal::Constant(c)) => self.visit(w, self.0.get(w).unwrap(), visited) | c,
                (Signal::Constant(c), Signal::Wire(w)) => c | self.visit(w, self.0.get(w).unwrap(), visited),
                (Signal::Constant(c1), Signal::Constant(c2)) => c1 | c2,
            },
            Cmd::Lshift(a,b) => match a {
                Signal::Wire(w) => self.visit(w, self.0.get(w).unwrap(), visited) << b,
                Signal::Constant(c) => c << b,
            },
            Cmd::Rshift(a,b) => match a {
                Signal::Wire(w) => self.visit(w, self.0.get(w).unwrap(), visited) >> b,
                Signal::Constant(c) => c >> b,
            }
        };
        visited.insert(k.to_string(), r);
        r
    }
}

fn p1(s: &str) -> u16 {
    *calc(s).get("a").unwrap()
}

fn p2(s: &str) -> u16 {
    let mut ops = s.lines().map(|v| v.parse::<Op>()).collect::<Result<Vec<_>, _>>().expect("parsing error");
    let mut g = Graph::new(ops.clone());

    let a = g.process().get("a").unwrap().clone();

    for o in &mut ops {
        if o.1 == "b" {
            o.0 = Cmd::Load(Signal::Constant(a));
            break;
        }
    }
    g = Graph::new(ops);

    *g.process().get("a").unwrap()
}
// Now, take the signal you got on wire a, 
// override wire b to that signal, and reset the other wires (including wire a). 
// What new signal is ultimately provided to wire a?


fn calc(s: &str) -> HashMap<String, u16> {
    let ops = s.lines().map(|v| v.parse::<Op>()).collect::<Result<Vec<_>, _>>().expect("parsing error");
    let g = Graph::new(ops);

    g.process()
}