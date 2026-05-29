package datamodeling

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func TestParseEvent(t *testing.T) {
	rows := []string{
		"u1, search,2026-05-20T10:00:00Z,120",
		"u2, checkout,2026-05-20T10:01:00Z,350",
	}

	got, err := ParseEvents(rows)
	if err != nil {
		t.Fatalf("Parse Events returned error : %v ", err)
	}
	if len(got) != 2 {
		t.Fatalf("len(events) = %d  want 2", len(got))
	}
}

func TestParseEventMalformed(t *testing.T) {
	_, err := ParseEvents([]string{"u1, search,not a time,120"})
	if err == nil {
		t.Fatalf("expected error for malformed timestamp")
	}

	if !strings.Contains(err.Error(), "row 1") {
		t.Fatalf("error should include row number, got : %v", err)
	}
}

func Test95Latency(t *testing.T) {
	events := []Event{
		{Action: "search", LatencyMS: 10},
		{Action: "search", LatencyMS: 20},
		{Action: "search", LatencyMS: 30},
		{Action: "search", LatencyMS: 40},
		{Action: "checkOut", LatencyMS: 999},
	}
	got, ok := P95Latency(events, "search")

	if !ok {
		t.Fatalf("expected matching search events")
	}

	if got != 40 {
		t.Fatalf("P95Latency == %d but want 40", got)
	}

}

func TestActionByUser(t *testing.T) {
	t1 := time.Date(2026, 5, 20, 10, 0, 0, 0, time.UTC)
	t2 := t1.Add(time.Minute)
	t3 := t1.Add(2 * time.Minute)

	events := []Event{
		{UserId: "u1", Action: "cehckout", Timestamp: t3},
		{UserId: "u2", Action: "search", Timestamp: t1},
		{UserId: "u1", Action: "search", Timestamp: t1},
		{UserId: "u1", Action: "add-to-cart", Timestamp: t2},
	}

	want := map[string][]string{
		"u1": {"search", "add-to-cart", "cehckout"},
		"u2": {"search"},
	}
	got := ActionByUser(events)

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("ActionByUser = %#v, want %#v", got, want)
	}

}
