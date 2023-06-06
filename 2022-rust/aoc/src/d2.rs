#![allow(dead_code, non_snake_case)]

const EXAMPLE_DATA: &str = "A Y
B X
C Z";

#[derive(Clone)]
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

struct WinInstruction(Outcome);
struct OpponentMove(Move);
struct MyMove(Move);

fn map_enemy(c: &str) -> Result<OpponentMove, String> {
    match c {
        "A" => Ok(OpponentMove(Move::Rock)),
        "B" => Ok(OpponentMove(Move::Paper)),
        "C" => Ok(OpponentMove(Move::Scissors)),
        _ => Err(format!("invalid enemy move '{}'", c))
    }
}

fn map_player(c: &str) -> Result<MyMove, String> {
    match c {
        "X" => Ok(MyMove(Move::Rock)),
        "Y" => Ok(MyMove(Move::Paper)),
        "Z" => Ok(MyMove(Move::Scissors)),
        _ => Err(format!("invalid player move '{}'", c))
    }
}

fn map_strategy(c: &str) -> Result<WinInstruction, String> {
    match c {
        "X" => Ok(WinInstruction(Outcome::Lose)),
        "Y" => Ok(WinInstruction(Outcome::Draw)),
        "Z" => Ok(WinInstruction(Outcome::Win)),
        _ => Err(format!("invalid strategy move '{}'", c))
    }
}

fn find_my_move(oppMove: &OpponentMove, strat: &WinInstruction) -> MyMove {
    match (&oppMove.0, &strat.0) {

        (Move::Rock, Outcome::Win) => MyMove(Move::Paper),
        (Move::Rock, Outcome::Lose) => MyMove(Move::Scissors),

        (Move::Paper, Outcome::Win) => MyMove(Move::Scissors),
        (Move::Paper, Outcome::Lose) => MyMove(Move::Rock),

        (Move::Scissors, Outcome::Win) => MyMove(Move::Rock),
        (Move::Scissors, Outcome::Lose) => MyMove(Move::Paper),

        (_, Outcome::Draw) => MyMove(oppMove.0.clone())
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
    assert_eq!(Ok(15), solve_p1(EXAMPLE_DATA));
    assert_eq!(Ok(14827), solve_p1(&std::fs::read_to_string("d2.txt").unwrap()));
}

#[test]
fn p2_test() {
    assert_eq!(Ok(12), solve_p2(EXAMPLE_DATA));
    assert_eq!(Ok(13889), solve_p2(&std::fs::read_to_string("d2.txt").unwrap()));
}

fn split(line: &str) -> Result<(&str, &str), String> {
    line
    .split_once(' ')
    .ok_or(format!("Invalid line: {}", line))
}

fn splt_line_p1(line: &str) -> Result<(OpponentMove, MyMove), String> {
    let raw = split(line)?;

    let enemy = map_enemy(raw.0)?;
    let my = map_player(raw.1)?;

    Ok((enemy, my))
}

fn splt_line_p2(line: &str) -> Result<(OpponentMove, WinInstruction), String> {
    let raw = split(line)?;

    let enemy = map_enemy(raw.0)?;
    let strat = map_strategy(raw.1)?;

    Ok((enemy, strat))
}

fn do_game(v: Result<(OpponentMove, MyMove), String>) -> Result<(MyMove, Outcome), String> {
    v.map(|vals| {
        let out = game(&vals.1, &vals.0);
        (vals.1, out)
    })
}

fn get_score(v: Result<(MyMove, Outcome), String>) -> Result<i32, String> {
    v.map(|vals|score(&vals.0, &vals.1))
}

fn solve_p1(data: &str) -> Result<i32, String> {
    let res: Result<Vec<i32>, String> = data.lines()
    .map(splt_line_p1)
    .map(do_game)
    .map(get_score)
    .collect();
    
    res.map(|v| v.iter().sum())
}

fn solve_p2(data: &str) -> Result<i32, String> {
    let res: Result<Vec<i32>, String> = data.lines()
    .map(splt_line_p2)
    .map(|v| v.map(|vals| {
        let myMove = find_my_move(&vals.0, &vals.1);
        (vals.0, myMove)
    }))
    .map(do_game)
    .map(get_score)
    .collect();

    res.map(|v| v.iter().sum())
}