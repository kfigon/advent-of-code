use std::{fs, collections::{HashMap}};

static EXAMPLE: &'static str = "1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet";

static EXAMPLE_2: &'static str = "two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen";

#[test]
fn p1_ex() {
    assert_eq!(142, p1(EXAMPLE));
}

#[test]
fn p2_ex() {
    assert_eq!(281, p2(EXAMPLE_2));
}

#[test]
fn p1_test() {
    assert_eq!(55607, p1(&fs::read_to_string("d1.txt").unwrap()));
}

#[test]
fn p2_test() {
    assert_eq!(55291, p2(&fs::read_to_string("d1.txt").unwrap()));
}

fn parse(s: &str, nums: impl Fn(&str) -> (Option<u32>, Option<u32>)) -> Vec<u32> {
    s.lines()
        .map(nums)
        .map(|pair| match pair {
            (Some(a), Some(b)) => Some(a*10 + b),
            _ => None,
        })
        .flatten()
        .collect()
}

fn p1(s: &str) -> u32 {
    let func = |line: &str| -> (Option<u32>, Option<u32>) {
        (
            line.chars().find(|c| c.is_ascii_digit()).and_then(|c| c.to_digit(10)),
            line.chars().rev().find(|c| c.is_ascii_digit()).and_then(|c| c.to_digit(10))
        )
    };

    parse(s, func).iter().sum()
}

fn p2(s: &str) -> u32 {
    let allowed_numbers: HashMap<&str, u32> = HashMap::from_iter(
        [
            ("one", 1),
            ("two", 2),
            ("three", 3),
            ("four", 4),
            ("five", 5),
            ("six", 6),
            ("seven", 7),
            ("eight", 8),
            ("nine", 9),
            ("zero", 0),
            ("1", 1),
            ("2", 2),
            ("3", 3),
            ("4", 4),
            ("5", 5),
            ("6", 6),
            ("7", 7),
            ("8", 8),
            ("9", 9),
            ("0", 0),
        ]
    );
    let func = |line: &str| -> (Option<u32>, Option<u32>) {
        let from_start = allowed_numbers.iter()
            .map(|(&pattern, &val)| line.find(pattern).map(|idx| (idx, val)))
            .flatten()
            .min_by(|a,b| a.0.cmp(&b.0));

        let from_end: Option<(usize, u32)> = allowed_numbers.iter()
            .map(|(&pattern, &val)| line.rfind(pattern).map(|idx| (idx, val)))
            .flatten()
            .max_by(|a,b| a.0.cmp(&b.0));

        match (from_start, from_end) {
            (Some(a), Some(b)) => (Some(a.1), Some(b.1)),
            _ => (None, None),
        }
    };

    parse(s, func).iter().sum()
}