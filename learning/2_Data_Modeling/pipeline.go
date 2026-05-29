package datamodeling

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Event struct {
	UserId    string
	Action    string
	Timestamp time.Time
	LatencyMS int
}

func ParseEvents(rows []string) ([]Event, error) {
	var events []Event
	for i, row := range rows {

		parts := strings.Split(row, ",")
		if len(parts) != 4 {
			return nil, fmt.Errorf("malformed row at index %d: %s", i+1, row)
		}

		userId := parts[0]
		action := parts[1]
		timestamp, err := time.Parse(time.RFC3339, parts[2])
		if err != nil {
			return nil, fmt.Errorf("malformed timestamp at index row %d: %s", i+1, parts[2])
		}
		latencyMS, err := strconv.Atoi(parts[3])
		if err != nil {
			return nil, fmt.Errorf("malformed latency at index %d: %s", i+1, parts[3])
		}
		events = append(events, Event{
			UserId:    userId,
			Action:    action,
			Timestamp: timestamp,
			LatencyMS: latencyMS,
		})

	}
	return events, nil
}

func P95Latency(events []Event, action string) (int, bool) {

	var latencies []int
	for _, event := range events {
		if event.Action == action {
			latencies = append(latencies, event.LatencyMS)
		}
		if len(latencies) == 0 {
			return 0, false
		}

	}
	sort.Ints(latencies)
	return latencies[len(latencies)*95/100], true
}

func ActionByUser(events []Event) map[string][]string {
	actionMap := make(map[string][]string)
	sort.Slice(events, func(i, j int) bool {
		return events[i].Timestamp.Before(events[j].Timestamp)
	})
	for _, event := range events {
		actionMap[event.UserId] = append(actionMap[event.UserId], event.Action)
	}
	return actionMap
}
