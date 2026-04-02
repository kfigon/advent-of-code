package d4

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var example = `[1518-11-01 00:00] Guard #10 begins shift
[1518-11-03 00:24] falls asleep
[1518-11-01 00:25] wakes up
[1518-11-01 00:55] wakes up
[1518-11-01 00:05] falls asleep
[1518-11-03 00:05] Guard #10 begins shift
[1518-11-01 23:58] Guard #99 begins shift
[1518-11-02 00:40] falls asleep
[1518-11-02 00:50] wakes up
[1518-11-04 00:36] falls asleep
[1518-11-03 00:29] wakes up
[1518-11-05 00:03] Guard #99 begins shift
[1518-11-04 00:02] Guard #99 begins shift
[1518-11-04 00:46] wakes up
[1518-11-01 00:30] falls asleep
[1518-11-05 00:45] falls asleep
[1518-11-05 00:55] wakes up`

func TestP1(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		got, err := p1(example)
		require.NoError(t, err)
		require.Equal(t, 240, got)
	})

	t.Run("real", func(t *testing.T) {
		data, err := os.ReadFile("d4.txt")
		require.NoError(t, err)

		got, err := p1(string(data))
		require.NoError(t, err)
		require.Equal(t, 39422, got)
	})
}

func TestP2(t *testing.T) {
	t.Run("ex", func(t *testing.T) {
		got, err := p2(example)
		require.NoError(t, err)
		require.Equal(t, 4455, got)
	})

	t.Run("real", func(t *testing.T) {
		data, err := os.ReadFile("d4.txt")
		require.NoError(t, err)

		got, err := p2(string(data))
		require.NoError(t, err)
		require.Equal(t, 65474, got)
	})
}
func p1(data string) (int, error) {
	expandedSleepTimes, err := sleepReport(data)
	if err != nil {
		return 0, err
	}

	var maxSleepGuardID *guardID
	for id, v := range expandedSleepTimes {
		if maxSleepGuardID == nil {
			maxSleepGuardID = &id
			continue
		}
		got := expandedSleepTimes[*maxSleepGuardID]
		if sumMinutes(v) > sumMinutes(got) {
			maxSleepGuardID = &id
		}
	}
	if maxSleepGuardID == nil {
		return 0, fmt.Errorf("sleep times not found")
	}

	sleptMinutes := expandedSleepTimes[*maxSleepGuardID]
	var maxMinute *int
	for m, v := range sleptMinutes {
		if maxMinute == nil {
			maxMinute = &m
			continue
		}
		got := sleptMinutes[*maxMinute]
		if v > got {
			maxMinute = &m
		}
	}

	return int(*maxSleepGuardID) * *maxMinute, nil
}

func p2(data string) (int, error) {
	expandedSleepTimes, err := sleepReport(data)
	if err != nil {
		return 0, err
	}

	var maxGuard *guardID
	var maxMinute *int
	for guard, minuteReport := range expandedSleepTimes {
		for minute, cnt := range minuteReport {
			if maxGuard == nil {
				maxGuard = &guard
			}
			if maxMinute == nil {
				maxMinute = &minute
			}

			got := expandedSleepTimes[*maxGuard][*maxMinute]
			if got < cnt {
				maxGuard = &guard
				maxMinute = &minute
			}
		}
	}

	if maxGuard == nil || maxMinute == nil {
		return 0, fmt.Errorf("result not found")
	}
	return *maxMinute * int(*maxGuard), nil
}

func sleepReport(data string) (map[guardID]map[int]int, error) {
	timeline, err := parseAndSort(data)
	if err != nil {
		return nil, err
	}
	expandedSleepTimes := map[guardID]map[int]int{}

	lastGuardID := -1
	lastSleepID := -1
	for i, ev := range timeline {
		switch ev.eventType {
		case begins:
			if ev.guardID == nil {
				return nil, fmt.Errorf("error on line %d, missing guard id", i)
			}
			lastGuardID = i
		case FallsAsleep:
			lastSleepID = i
		case WakesUp:
			if lastGuardID == -1 || lastSleepID == -1 {
				return nil, fmt.Errorf("corrupted data, guard or sleep not found for wakeup, idx: %d", i)
			}

			min, err := calcSleepMinutes(timeline[lastSleepID].time, timeline[i].time)
			if err != nil {
				return nil, fmt.Errorf("error calculating sleep minutes on id %d: %w", i, err)
			}

			for _, m := range min {
				_, ok := expandedSleepTimes[*timeline[lastGuardID].guardID]
				if !ok {
					expandedSleepTimes[*timeline[lastGuardID].guardID] = map[int]int{}
				}
				expandedSleepTimes[*timeline[lastGuardID].guardID][m]++
			}
		default:
			return nil, fmt.Errorf("invalid event: %v", ev.eventType)
		}
	}
	return expandedSleepTimes, nil
}

func sumMinutes(m map[int]int) int {
	out := 0
	for _, v := range m {
		out += v
	}
	return out
}

// it's always single night, no need to check a lot of fields. It's always the minute part
func calcSleepMinutes(before, after time) ([]int, error) {
	if sortTime(before, after) > 0 {
		return nil, fmt.Errorf("items are not sorted, %v should be before %v", before, after)
	}

	out := make([]int, 0, after.minute-before.minute)
	for i := before.minute; i < after.minute; i++ {
		out = append(out, i)
	}
	return out, nil
}

func parseAndSort(data string) ([]event, error) {
	timeline := []event{}
	for line := range strings.Lines(data) {
		line = strings.TrimSpace(line)
		e, err := parse(line)
		if err != nil {
			return nil, err
		}
		timeline = append(timeline, e)
	}

	slices.SortFunc(timeline, func(a, b event) int {
		return sortTime(a.time, b.time)
	})
	return timeline, nil
}

type eventType int

const (
	begins eventType = iota
	FallsAsleep
	WakesUp
)

type time struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
}

type guardID int
type event struct {
	guardID   *guardID // only valid when begins
	eventType eventType
	time      time
}

func sortTime(at, bt time) int {
	if at.year < bt.year {
		return -1
	} else if at.year > bt.year {
		return 1
	} else if at.month < bt.month {
		return -1
	} else if at.month > bt.month {
		return 1
	} else if at.day < bt.day {
		return -1
	} else if at.day > bt.day {
		return 1
	} else if at.hour < bt.hour {
		return -1
	} else if at.hour > bt.hour {
		return 1
	} else if at.minute < bt.minute {
		return -1
	} else if at.minute > bt.minute {
		return 1
	}
	return 0
}

func parse(line string) (event, error) {
	parts := strings.FieldsFunc(line, func(r rune) bool {
		return map[rune]bool{
			'[': true,
			']': true,
			'-': true,
			':': true,
			'#': true,
			' ': true,
		}[r]
	})
	var zero event
	if len(parts) < 6 {
		return zero, fmt.Errorf("cant parse event: %s", line)
	}

	year, err := strconv.Atoi(parts[0])
	if err != nil {
		return zero, fmt.Errorf("error parsing year: %s: %w", line, err)
	}
	month, err := strconv.Atoi(parts[1])
	if err != nil {
		return zero, fmt.Errorf("error parsing month: %s: %w", line, err)
	}
	day, err := strconv.Atoi(parts[2])
	if err != nil {
		return zero, fmt.Errorf("error parsing day: %s: %w", line, err)
	}
	hour, err := strconv.Atoi(parts[3])
	if err != nil {
		return zero, fmt.Errorf("error parsing hour: %s: %w", line, err)
	}
	minute, err := strconv.Atoi(parts[4])
	if err != nil {
		return zero, fmt.Errorf("error parsing minute: %s: %w", line, err)
	}

	var kind *eventType
	var guard *guardID
	p := parts[6]

	if p == "asleep" {
		kind = ptr(FallsAsleep)
	} else if p == "up" {
		kind = ptr(WakesUp)
	} else if v, err := strconv.Atoi(p); err != nil {
		return zero, fmt.Errorf("error parsing guard id for line: %s, checked field: %s, err: %w", line, p, err)
	} else {
		guard = ptr(guardID(v))
		kind = ptr(begins)
	}

	return event{
		guardID:   guard,
		eventType: *kind,
		time: time{
			year:   year,
			month:  month,
			day:    day,
			hour:   hour,
			minute: minute,
		},
	}, nil
}

func ptr[T any](v T) *T {
	return &v
}
