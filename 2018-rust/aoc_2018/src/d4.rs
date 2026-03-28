use std::{cmp::Ordering, collections::HashMap, str::FromStr};

#[derive(Debug)]
struct Time {
    year: usize,
    month: usize,
    day: usize,
    hour: usize,
    minute: usize,
}

#[derive(Debug)]
struct Event {
    time: Time,
    event: EventType,
}

impl FromStr for Event {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        // [1518-11-03 00:05] Guard #10 begins shift

        let parts = s.split(&['[', '-', ']', ' ', '#'][..])
            .filter(|c| !c.is_empty())
            .collect::<Vec<&str>>();

        let year:usize = parts.get(0).ok_or(format!("year not found in {}", s))?.parse().map_err(|err| format!("year parsing error {} for {}", err, s))?;
        let month:usize = parts.get(1).ok_or(format!("month not found in {}", s))?.parse().map_err(|err| format!("month parsing error {} for {}", err, s))?;
        let day:usize = parts.get(2).ok_or(format!("day not found in {}", s))?.parse().map_err(|err| format!("day parsing error {} for {}", err, s))?;
        let time:(usize, usize) = parse_time(parts.get(3).ok_or(format!("time not found in {}", s))?)?;
        let event:EventType = parts.get(5).ok_or(format!("event not found in {}", s))?.parse().map_err(|err| format!("event parsing error {} for {}", err, s))?;

        let hour = time.0;
        let minute = time.1;
        Ok(Event{ time: Time { year, month, day, hour, minute, },event })
    }
}

fn parse_time(s: &str) -> Result<(usize, usize), String> {
    match s.split_once(":") {
        Some((hour, minute)) => match (hour.parse::<usize>(), minute.parse::<usize>()) {
            (Ok(hourV), Ok(minuteV)) => Ok((hourV, minuteV)),
            _ => Err(format!("time parsing error for {}", s))
        }
        _ => Err(format!("time not found for {}", s))
    }
}

type ID = usize;

#[derive(Debug)]
enum EventType {
    Begin(ID),
    WakeUp,
    Sleep,
}
impl FromStr for EventType {
    type Err = String;

    fn from_str(s: &str) -> Result<Self, Self::Err> {
        if s.contains("up") {
            return Ok(Self::WakeUp);
        } else if s.contains("asleep") {
            return Ok(Self::Sleep);
        } else if let Ok(num) = s.parse::<usize>() {
            return Ok(Self::Begin(num));
        }
        Err(format!("invalid event: {}", s))
    }
}

type Timeline = Vec<Event>;

fn p1(input: &str) -> Result<usize, String> {
    let timeline = parse_input(input)?;
    let longest = longest_sleep(&timeline);
    match longest {
        Some((id, len)) => Ok(id*len),
        None => Err("not found".to_string()),
    }
}

fn parse_input(input: &str) -> Result<Timeline, String> {
    // todo: parse and sort
    let mut timeline: Vec<Event> = input.lines()
        .map(&str::parse::<Event>)
        .collect::<Result<Timeline, String>>()?;
    timeline.sort_by(|a,b| {
        let at = &a.time;
        let bt = &b.time;
        if at.year < bt.year {
            return Ordering::Less;
        } else if at.year > bt.year {
            return Ordering::Greater;
        } else if at.month < bt.month{
            return Ordering::Less;
        } else if at.month > bt.month {
            return Ordering::Greater;
        } else if at.day < bt.day {
            return Ordering::Less;
        } else if at.day > bt.day {
            return Ordering::Greater;
        } else if at.hour < bt.hour {
            return Ordering::Less;
        } else if at.hour > bt.hour {
           return Ordering::Greater; 
        } else if at.minute < bt.minute {
            return Ordering::Less;
        } else if at.minute > bt.minute{
            return Ordering::Greater;
        }
        return Ordering::Equal;
    });
    Ok(timeline)
}

fn longest_sleep(timeline: &Timeline) -> Option<(ID, usize)> {
    let mut sleep_times = HashMap::<ID, usize>::new();
    enum State {
        start,
        guardEvent(ID),
        fellAsleep(ID, Time),
        wokeUp(ID, Time),
    }
    
    
    let mut parser_state = State::start;
    for t in timeline {
        match (&t.event, parser_state) {
            (EventType::Begin(new_id), State::start) => {
                parser_state = State::guardEvent(*new_id);
            }
            (EventType::Begin(new_id), State::guardEvent(id)) => todo!(),
            (EventType::Begin(new_id), State::fellAsleep(id, time)) => todo!(),
            (EventType::Begin(new_id), State::wokeUp(id, time)) => todo!(),
            (EventType::WakeUp, State::guardEvent(id)) => todo!(),
            (EventType::WakeUp, State::fellAsleep(id, time)) => todo!(),
            (EventType::WakeUp, State::wokeUp(id, time)) => todo!(),
            (EventType::Sleep, State::guardEvent(id)) => todo!(),
            (EventType::Sleep, State::fellAsleep(id, time)) => todo!(),
            (EventType::Sleep, State::wokeUp(id, time)) => todo!(),

            (EventType::Sleep, State::start) => unreachable!(),
            (EventType::WakeUp, State::start) => unreachable!(),
        }
    }

    sleep_times.iter().max()
}

#[cfg(test)]
mod test {
    use super::*;

    #[test]
    fn p1_ex() {
        let data = "[1518-11-03 00:05] Guard #10 begins shift
[1518-11-03 00:29] wakes up
[1518-11-02 00:40] falls asleep
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-02 00:50] wakes up
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-04 00:46] wakes up
[1518-11-04 00:36] falls asleep
[1518-11-01 00:05] falls asleep
[1518-11-03 00:24] falls asleep
[1518-11-05 00:45] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-01 00:55] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-05 00:55] wakes up";


        let expected = 240;
        assert_eq!(Ok(expected), p1(data));
    }
}