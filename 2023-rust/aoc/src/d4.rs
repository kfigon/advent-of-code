use std::{collections::HashSet, str::FromStr, fs};

static EXAMPLE: &'static str= "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11";

#[test]
fn p1_ex() {
    assert_eq!(13, p1(EXAMPLE));
}

#[test]
fn p1_test() {
    assert_eq!(23847, p1(&fs::read_to_string("d4.txt").unwrap()));
}

struct Card {
    id: usize,
    winning: HashSet<usize>,
    guesses: HashSet<usize>,
}

impl FromStr for Card {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        let card = s.split_once(": ").ok_or(format!("no card id found in {s}"))?;
        let card_num = card.0.split_once(" ")
            .ok_or(format!("no id found in {s}"))?
            .1.trim().parse::<usize>().map_err(|e| format!("{e}: Cant parse card num in {s}"))?;

        let nums = card.1.split_once(" | ").ok_or(format!("cant find winning and guesses in {s}"))?;
        let winning = nums.0.split_whitespace()
            .map(|v| v.parse::<usize>())
            .collect::<Result<HashSet<usize>, _>>()
            .map_err(|e| format!("{e}: cant parse winning numbers {s}"))?;

        let guesses = nums.1.split_whitespace()
            .map(|v| v.parse::<usize>())
            .collect::<Result<HashSet<usize>, _>>()
            .map_err(|e| format!("{e}: cant parse guessed numbers {s}"))?;

        Ok(Card { id: card_num, winning: winning, guesses: guesses })
    }
}

fn p1(s: &str) -> usize {
    let cards = s.lines()
        .map(|line| line.parse::<Card>())
        .collect::<Result<Vec<Card>, _>>()
        .unwrap();

    cards.into_iter()
        .map(|c| c.winning.intersection(&c.guesses).count())
        .map(|num_of_winning| match num_of_winning {
            0 => 0,
            n => 2usize.pow(n as u32 -1),
        })
        .sum()
}