package main

import (
	"fmt"
	"strings"
)

const cutSet string = ",.?;:!\"'/"
const txt string = "As far as eye could reach he saw nothing but the stems of the great plants about him receding in the violet shade, and far overhead the multiple transparency of huge leaves filtering the sunshine to the solemn splendour of twilight in which he walked. Whenever he felt able he ran again; the ground continued soft and springy, covered with the same resilient weed which was the first thing his hands had touched in Malacandra. Once or twice a small red creature scuttled across his path, but otherwise there seemed to be no life stirring in the wood; nothing to fear—except the fact of wandering unprovisioned and alone in a forest of unknown vegetation thousands or millions of miles beyond the reach or knowledge of man."

func cntFreq(txt string) map[string]int {
	words := strings.Fields(txt)
	trimmedWords := make([]string, 0)
	for _, word := range words {
		trimmedWords = append(trimmedWords, strings.ToLower(strings.Trim(word, cutSet)))
	}

	freq := make(map[string]int, 10)
	for _, word := range trimmedWords {
		freq[word]++
	}
	return freq
}
func main() {
	freq := cntFreq(txt)
	for word, f := range freq {
		if f > 1 {
			fmt.Printf("%v: %v\n", word, f)
		}
	}
}
