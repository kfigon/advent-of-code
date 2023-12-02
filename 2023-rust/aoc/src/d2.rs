use std::fs;

static EXAMPLE: &'static str = "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green";

#[test]
fn p1_ex() {
    assert_eq!(8, p1(EXAMPLE));
}

#[test]
fn p1_test() {
    assert_eq!(3059, p1(&fs::read_to_string("d2.txt").unwrap()));
}

#[test]
fn p2_ex() {
    assert_eq!(2286, p2(EXAMPLE));
}

#[test]
fn p2_test() {
    assert_eq!(65371, p2(&fs::read_to_string("d2.txt").unwrap()));
}


struct Guess{
    red: usize, 
    green: usize, 
    blue: usize
}

impl TryFrom<&str> for Guess {
    type Error = String;

    fn try_from(value: &str) -> Result<Self, Self::Error> {
        let mut green: Option<usize> = None;
        let mut red: Option<usize> = None;
        let mut blue: Option<usize> = None;
        let set_v = |v: usize, target: &mut Option<usize>| -> Result<(), usize> {
            if let Some(existing) = target {
                return Err(*existing);
            }
            *target = Some(v);
            Ok(())
        };

        for p in value.split(", ") {
            let val = p.split_ascii_whitespace().collect::<Vec<_>>();
            if val.len() != 2 {
                return Err(format!("invalid input, expected pairs of num and color: {}", value));
            }
            match (val[0].parse::<usize>(), val[1]) {
                (Ok(num), "red") => set_v(num, &mut red).map_err(|e| format!("red set twice with {e} on: {value}"))?,
                (Ok(num), "green") => set_v(num, &mut green).map_err(|e| format!("green set twice with {e} on: {value}"))?,
                (Ok(num), "blue") =>  set_v(num, &mut blue).map_err(|e| format!("blue set twice with {e} on: {value}"))?,
                (Ok(_), c) =>  return Err(format!("invalid color {c}")),
                _ =>  return Err(format!("invalid input, no number: {value}")),
            };
        }

        Ok(Self { red: red.unwrap_or_default(), green: green.unwrap_or_default(), blue: blue.unwrap_or_default() })
    }
}

struct Game(usize, Vec<Guess>);

impl TryFrom<&str> for Game {
    type Error = String;

    fn try_from(value: &str) -> Result<Self, Self::Error> {
        let (game_num, guesses) = value.split_once(": ").ok_or(format!("cant find game delimiter {value}"))?;
        let guesses = guesses.split("; ");
        let guesses: Vec<Guess> = guesses.into_iter().map(|v| v.try_into()).collect::<Result<Vec<Guess>, _>>()?;

        let (_, game_num) = game_num
            .split_once(" ")
            .ok_or(format!("cant find game number {value}"))?;
        let game_num = game_num.parse::<usize>().map_err(|_| format!("invalid game number: {value}"))?;

        Ok(Game(game_num, guesses))
    }
}

impl Game {
    fn valid(&self) -> bool {
        self.1
            .iter()
            .all(|g| g.red <= 12 && g.green <= 13 && g.blue <= 14)
    }

    fn fewest_cubes(&self) -> Guess {
        let red = self.1.iter().map(|v| v.red).max().unwrap_or_default();
        let green = self.1.iter().map(|v| v.green).max().unwrap_or_default();
        let blue = self.1.iter().map(|v| v.blue).max().unwrap_or_default();

        Guess { red, green, blue }
    }
}

fn p1(s: &str) -> usize {
    let v = s.lines().map(|line| line.try_into()).collect::<Result<Vec<Game>, _>>().unwrap();

    v.iter()
        .filter(|g| g.valid())
        .map(|g| g.0)
        .sum()
}

fn p2(s: &str) -> usize {
    let v = s.lines().map(|line| line.try_into()).collect::<Result<Vec<Game>, _>>().unwrap();

    v.iter()
        .map(|g| g.fewest_cubes())
        .map(|g| g.red * g.green * g.blue)
        .sum()
}