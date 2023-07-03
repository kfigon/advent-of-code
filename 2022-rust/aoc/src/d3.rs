use std::collections::{HashMap, HashSet};

const example: &str = "vJrwpWtwJgWrhcsFMMfFFhFp
jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL
PmmdzqPrVvPwwTWBwg
wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn
ttgJtRGJQctTZtZT
CrZsJsPPZsGzwwsLwLmpwMDw";

#[test]
fn p1_ex() {
    assert_eq!(157, p1(example))
}

#[test]
fn p1_file() {
    assert_eq!(8493, p1(&std::fs::read_to_string("d3.txt").unwrap()))
}

#[test]
fn p2_ex() {
    assert_eq!(70, p2(example))
}

#[test]
fn p2_file() {
    assert_eq!(2552, p2(&std::fs::read_to_string("d3.txt").unwrap()))
}

fn get_scores() -> HashMap<char, usize> {
    ('a'..='z')
    .chain('A'..='Z')
    .enumerate()
    .map(|(idx, c)| (c, idx+1))
    .collect::<HashMap<char, usize>>()
}

fn p1(input: &str) -> i32 {
    let scores = get_scores();

    // let mut res: i32 = 0;
    // for line in input.lines() {
    //     let middle = line.len()/2;
    //     let part_a = &line[..middle].chars().collect::<HashSet<char>>();
    //     for c in &line[middle..].chars().collect::<HashSet<char>>() {
    //         if part_a.contains(&c) {
    //             res += *scores.get(&c).unwrap() as i32;
    //         }
    //     }
    // }
    // res

    input.lines()
        .map(|line| {
            let middle = line.len()/2;
            let part_a = &line[..middle].chars().collect::<HashSet<char>>();
            let part_b = &line[middle..].chars().collect::<HashSet<char>>();

            part_b.iter()
            .filter(|c| part_a.contains(&c))
            .map(|c| *scores.get(c).unwrap() as i32)
            .sum::<i32>()
        })
        .sum::<i32>()
}

fn p2(input: &str) -> i32 {
    let scores = get_scores();

    const limit: usize = 3;

    let mut res: i32 = 0;
    for group in input.lines().collect::<Vec<&str>>().chunks(limit) {
        let mut counts: HashMap<char, usize> = HashMap::new();
        group.iter()
            .for_each(|g| g.chars().collect::<HashSet<char>>().iter()
                .for_each(|v| *counts.entry(*v).or_insert(0) += 1)
            );

        res += counts.iter().filter(|v| *v.1 == limit)
            .map(|v| *scores.get(v.0).unwrap() as i32)
            .sum::<i32>();
    }

    res


    // let mut res: i32 = 0;
    // let lines: Vec<&str> = input.lines().collect();
    
    // for i in (0..lines.len()).step_by(3) {
    //     let elf1 = lines.get(i).unwrap().chars().collect::<HashSet<char>>();
    //     let elf2 = lines.get(i+1).unwrap().chars().collect::<HashSet<char>>();
    //     let elf3 = lines.get(i+2).unwrap().chars().collect::<HashSet<char>>();

    //     let mut counts: HashMap<char, i32> = HashMap::new();
    //     elf1.iter().for_each(|c| *counts.entry(*c).or_insert(0) += 1 );
    //     elf2.iter().for_each(|c| *counts.entry(*c).or_insert(0) += 1 );
    //     elf3.iter().for_each(|c| *counts.entry(*c).or_insert(0) += 1 );

    //     res += counts.iter()
    //         .filter(|v| *v.1 == 3)
    //         .map(|v| *scores.get(v.0).unwrap() as i32)
    //         .sum::<i32>();
    // }
    // res
}