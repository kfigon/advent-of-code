use std::os;

fn main() {
    println!("Hello, world!");
}

const example_data: &'static str = "A Y
B X
C Z";

enum Move {
    Rock,
    Paper,
    Scissors
}

enum Outcome {
    Win,
    Lose,
    Draw
}

fn map_enemy(c: &str) -> Result<Move, String> {
    match c {
        "A" => Ok(Move::Rock),
        "B" => Ok(Move::Paper),
        "C" => Ok(Move::Scissors),
        _ => Err(format!("invalid enemy move '{}'", c))
    }
}

fn map_player(c: &str) -> Result<Move, String> {
    match c {
        "X" => Ok(Move::Rock),
        "Y" => Ok(Move::Paper),
        "Z" => Ok(Move::Scissors),
        _ => Err(format!("invalid player move '{}'", c))
    }
}

fn game(myMove: &MyMove, enemyMove: &OpponentMove) -> Outcome {
    match (&myMove.0, &enemyMove.0) {
        (Move::Rock, Move::Rock) => Outcome::Draw,
        (Move::Paper, Move::Paper) => Outcome::Draw,
        (Move::Scissors, Move::Scissors) => Outcome::Draw,

        (Move::Rock, Move::Paper) => Outcome::Lose,
        (Move::Paper, Move::Scissors) => Outcome::Lose,
        (Move::Scissors, Move::Rock) => Outcome::Lose,

        (Move::Rock, Move::Scissors) => Outcome::Win,
        (Move::Paper, Move::Rock) => Outcome::Win,
        (Move::Scissors, Move::Paper) => Outcome::Win,
    }
}

fn score(myMove: &MyMove, outcome: &Outcome) -> i32 {
    let outcomeScore = match &outcome {
        Outcome::Win => 6,
        Outcome::Lose => 0,
        Outcome::Draw => 3,
    };

    let moveScore = match &myMove.0 {
        Move::Rock => 1,
        Move::Paper => 2,
        Move::Scissors => 3,
    };

    moveScore + outcomeScore
}

#[test]
fn p1_test() {
    assert_eq!(15, solve_p1(example_data));
    assert_eq!(14827, solve_p1(&std::fs::read_to_string("data.txt").unwrap()));
}

#[test]
fn p2_test() {
    assert_eq!(12, solve_p2(example_data));
    assert_eq!(13889, solve_p2(&std::fs::read_to_string("data.txt").unwrap()));
}

struct OpponentMove(Move);
struct MyMove(Move);

fn splt_line(line: &str) -> Result<(OpponentMove, MyMove), String> {
    let raw: (&str, &str) = line
    .split_once(" ")
    .ok_or(format!("Invalid line: {}", line))?;

    let enemy = map_enemy(raw.0)?;
    let my = map_player(raw.1)?;

    Ok((OpponentMove(enemy), MyMove(my)))
}

fn solve_p1(data: &str) -> i32 {
    let res: Result<Vec<i32>, String> = data.lines()
    .map(splt_line)
    .map(|v| v.map(|vals| {
        let out = game(&vals.1, &vals.0);
        (vals.1, out)
    }))
    .map(|v| v.map(|vals|score(&vals.0, &vals.1)))
    .collect();
    
    match res {
        Ok(v) => v.iter().sum(),
        Err(e) => {
            println!("{}", e);
            0
        }
    }
}

fn solve_p2(data: &str) -> i32 {
    todo!()
}