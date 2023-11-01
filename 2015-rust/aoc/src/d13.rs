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
        if res.len() != 5 {
            return Err("invalid input, parts not found");
        } 
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
}