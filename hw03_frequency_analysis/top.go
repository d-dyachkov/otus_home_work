package hw03_frequency_analysis //nolint:golint,stylecheck

import (
	"regexp"
	"sort"
	"strings"
)

type wordFrequency struct {
	word      string
	frequency int
}

var wordSplittingPattern = regexp.MustCompile(`[\s!.,?]+`)

func Top10(text string) []string {
	countionWords := analyzeWordFrequency(text)
	sort.Slice(countionWords, func(i, j int) bool { return countionWords[i].frequency > countionWords[j].frequency })
	return takeWords(countionWords, 10)
}

func analyzeWordFrequency(text string) []wordFrequency {
	wordCount := countingWords(text)
	words := make([]wordFrequency, 0, len(wordCount))
	for word, count := range wordCount {
		words = append(words, wordFrequency{word: word, frequency: count})
	}
	return words
}

func countingWords(text string) map[string]int {
	wordCount := make(map[string]int)
	for _, word := range wordSplittingPattern.Split(text, -1) {
		if word == "" || word == "-" {
			continue
		}
		word = strings.ToLower(word)
		wordCount[word]++
	}
	return wordCount
}

func takeWords(words []wordFrequency, n int) []string {
	if n > len(words) {
		n = len(words)
	}
	result := make([]string, 0, n)
	for _, word := range words[:n] {
		result = append(result, word.word)
	}
	return result
}
