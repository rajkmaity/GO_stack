package fundamentals

import (
	"sort"
	"strings"
)

//  Learn abou for-loop , array

func NormalizeScores(scores []float64) []float64 {
	// If the scores is empty then return []
	if len(scores) == 0 {
		return []float64{}
	}
	// Compure min and max
	min := scores[0]
	max := scores[0]

	for _, i := range scores {
		if min > i {
			min = i
		}
		if max < i {
			max = i
		}
	}

	if min == max {
		return make([]float64, len(scores))
	}

	results := make([]float64, len(scores))

	for i, score := range scores {
		results[i] = (score - min) / (max - min)
	}
	return results

}

// String Manipulation
// Sorting

func TopKWords(text string, k int) []string {
	// Check the corner case
	if text == "" || k <= 0 {
		return []string{}
	}
	var words []string
	wordCount := make(map[string]int)
	var currentWord string

	for _, ch := range text {
		if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			currentWord += string(ch)
		} else {
			if currentWord != "" {
				word := strings.ToLower(currentWord)
				if wordCount[word] == 0 {
					words = append(words, word)
				}
				wordCount[word]++
				currentWord = ""
			}
		}
	}
	if currentWord != "" {
		word := strings.ToLower(currentWord)
		if wordCount[word] == 0 {
			words = append(words, word)
		}
		wordCount[word]++
	}

	sort.Slice(words, func(i, j int) bool {
		if wordCount[words[i]] != wordCount[words[j]] {
			return wordCount[words[i]] > wordCount[words[j]]
		}
		return words[i] < words[j]
	})

	if k > len(words) {
		k = len(words)
	}
	return words[:k]
}

/*
func main() {
	text := "go go do do Python Python"
	fmt.Println(TopKWords(text, 3))
}
*/
