use std::vec;

const input: &str = "    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3 

move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2";

#[derive(Debug, PartialEq)]
struct Stack<T>(Vec<T>);

#[derive(Debug, PartialEq)]
struct Cmd<'a>(&'a str, &'a str, &'a str);

struct Input<'a>{
    cranes: Vec<Stack<&'a str>>,
    commands: Vec<Cmd<'a>>
}

impl<'a> Input<'_> {
    fn parse(inp: &'a str) -> Result<Input, &str> {
        let parts = inp.split("\n\n").collect::<Vec<&str>>();
        if parts.len() != 2 {
            return Err("invalid input");
        }
        let cranes_raw = *parts.get(0).unwrap();
        let commands_raw = *parts.get(1).unwrap();

        let cranes_lines: Vec<_> = cranes_raw.lines().map(map_crane_line).collect();
        let cranes = todo!();

        Ok(Input { 
                cranes: cranes, 
                commands: commands_raw.lines().filter_map(map_command_line).collect(),
            }
        )
    }
}

fn map_crane_line(line: &str) -> Vec<Option<String>> {
    line.chars()
        .collect::<Vec<char>>()
        .chunks(4)
        .map(|v| {
            let st = v.iter().collect::<String>();
            let s = st.trim();
            if s.is_empty() {
                None
            } else {
                Some(s.to_string())
            }
        }).collect::<Vec<Option<String>>>()
    }

fn map_command_line(line: &str) -> Option<Cmd> {
    // move 1 from 2 to 1
    let parts = line.split(" from ").collect::<Vec<_>>();

    let what = parts.get(0).map(|v| v.split("move ").into_iter()).and_then(|mut v| v.nth(1))?;
    let from = parts.get(1).map(|v| v.split(" to ").into_iter()).and_then(|mut v| v.nth(0))?;
    let to = parts.get(1).map(|v| v.split(" to ").into_iter()).and_then(|mut v| v.nth(1))?;
    
    Some(Cmd(what,from,to))
}

#[test]
fn parse_example() {
    let i = Input::parse(input).expect("expected valid input");
    assert_eq!(i.cranes, vec![
        Stack(vec!["N","Z"]),
        Stack(vec!["D","C","M"]),
        Stack(vec!["P"])],
    );

    assert_eq!(i.commands, vec![
        Cmd("1","2","1"),
        Cmd("3","1","3"),
        Cmd("2","2","1"),
        Cmd("1","1","2"),
    ]);
}

#[test]
fn parse_crane() {
    let r = map_crane_line("[N] [C]        ");
    assert_eq!(r, vec![Some("[N]".to_string()), Some("[C]".to_string()), None, None])
}

#[test]
fn parse_command() {
    let r = map_command_line("move 1 from 4 to 5");
    assert_eq!(r, Some(Cmd("1", "4", "5")))
}