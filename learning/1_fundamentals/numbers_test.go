package fundamentals

import (
	"reflect"
	"testing"
)

func TestNormalizeScores(t *testing.T) {
	tests := []struct {
		name   string
		scores []float64
		want   []float64
	}{
		{name: "empty", scores: []float64{}, want: []float64{}},
		{name: "equal values", scores: []float64{10, 10, 10}, want: []float64{0, 0, 0}},
		{name: "mixed values", scores: []float64{10, 20, 30}, want: []float64{0, 0.5, 1}},
		{name: "mixed values", scores: []float64{-5, 0, +5}, want: []float64{0, 0.5, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NormalizeScores(tt.scores)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("NormalizeScores(%v) = %#v, want %#v", tt.scores, got, tt.want)
			}
		})
	}
}

func TestTopKWords(t *testing.T) {
	tests := []struct {
		name string
		text string
		k    int
		want []string
	}{
		{name: "", text: "", k: 3, want: []string{}},
		{name: "zero k", text: "go go Python", k: 0, want: []string{}},
		{name: "basic ", text: "Go go Python! Python, GO", k: 2, want: []string{"go", "python"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TopKWords(tt.text, tt.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("TopKWords(%q, %d) = %#v, want %#v", tt.text, tt.k, got, tt.want)
			}
		})
	}

}
