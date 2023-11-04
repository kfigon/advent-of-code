use std::{fs, collections::HashMap};

use once_cell::sync::Lazy;
use regex::Regex;

const EXAMPLE: &str = "Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.";

#[derive(Debug)]
struct Entry<'a> {
    name: &'a str,
    speed: usize,
    run_duration: usize,
    rest_time: usize
}

impl<'a> Entry<'a> {
    fn distance(&self, seconds: usize) -> usize {
        let cycle_time = self.run_duration + self.rest_time;
        let cycle_distance = self.run_duration * self.speed;
        let number_of_full_cycles = seconds/cycle_time;
        let remainder_seconds = seconds - number_of_full_cycles * cycle_time;
        let remainder_distance = std::cmp::min(remainder_seconds, self.run_duration)*self.speed;

        number_of_full_cycles*cycle_distance + remainder_distance
    }
}

impl<'a> TryFrom<&'a str> for Entry<'a> {
    type Error = &'static str;

    fn try_from(s: &'a str) -> Result<Self, Self::Error> {
        static RE: Lazy<Regex> = Lazy::new(|| Regex::new(r"(\w+) can fly (\d+) km/s for (\d+) seconds, but then must rest for (\d+) seconds").unwrap());
        
        let res = RE.captures(s).ok_or("invalid input, parts not found")?;
        match (
            res.get(1).map(|v| v.as_str()),
            res.get(2).map(|v| v.as_str()).map(|v| v.parse::<usize>()),
            res.get(3).map(|v| v.as_str()).map(|v| v.parse::<usize>()),
            res.get(4).map(|v| v.as_str()).map(|v| v.parse::<usize>()),
        ) {
            (Some(who), Some(Ok(speed)), Some(Ok(time)), Some(Ok(rest))) => Ok(Self{
                name: who,
                speed,
                run_duration: time,
                rest_time: rest,
            }),
            _ => Err("missing parts or invalid format"),
        }
    }
}

fn parse(s: &str) -> Result<Vec<Entry>, &'static str> {
    s.lines()
        .map(|v| v.try_into())
        .collect()
}

#[test]
fn p1_ex() {
    assert_eq!(1120, p1(parse(EXAMPLE).unwrap(), 1000))
}

#[test]
fn p2_ex() {
    assert_eq!(689, p2(parse(EXAMPLE).unwrap(), 1000))
}

#[test]
fn p1_test() {
    let raw_data = &fs::read_to_string("d14.txt").unwrap();
    assert_eq!(2660, p1(parse(raw_data).unwrap(), 2503)); 
}

#[test]
fn p2_test() {
    let raw_data = &fs::read_to_string("d14.txt").unwrap();
    assert_eq!(1256, p2(parse(raw_data).unwrap(), 2503));
}

fn p1(vs: Vec<Entry>, seconds: usize) -> usize {
    vs.iter().map(|v| v.distance(seconds)).max().unwrap_or_default()
}

fn p2(vs: Vec<Entry>, seconds: usize) -> usize {
    let participants = vs.iter().map(|v| v.name).collect::<Vec<_>>();
    let mut points: HashMap<&str, usize> = HashMap::from_iter(participants.iter().map(|v| (*v, 0)));

    for i in 1..=seconds {
        let distance: HashMap<&str, usize> = HashMap::from_iter(
            vs.iter().map(|v| (v.name, v.distance(i)))
        );

        let max_val = *distance.values().max().unwrap();
        // there can be multiple leaders, all needs to get a point
        distance.into_iter()
            .filter(|v| v.1 == max_val)
            .for_each(|leader| {
                *points.entry(leader.0).or_default() += 1;
            })
    }
    
    points.into_values().max().unwrap_or_default()
}