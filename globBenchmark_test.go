package glob

import (
	"bufio"
	"os"
	"regexp"
	"strings"
	"testing"
)

const wordsCount = 31207
const logLinesCount = 27992

var globPatterns = []string{
	"*a*c*",
	"*class*l*",
}

var greedyRegexPatterns = []string{
	"(?i)^.*a.*c.*$",
	"(?i)^.*class.*l.*$",
}

var nonGreedyRegexPatterns = []string{
	"(?i)^.*?a.*?c.*$",
	"(?i)^.*?class.*?l.*$",
}

var words []string
var logLines []string

func init() {
	words = readBenchmarkFile("words")
	logLines = readBenchmarkFile("logLines")
}

func BenchmarkGlogEqWords(b *testing.B) {

	word := words[100]
	m, err := Compile(word, WithCaseInsensitive(true))
	if err != nil {
		panic(err)
	}

	// Run the test
	for n := 0; n < b.N; n++ {
		var count int
		for _, w := range words {
			if m.Matches(w) {
				count += 1
			}
		}
		if count != 1 {
			b.Error("Failed to find the match")
		}
	}
}

func BenchmarkStringEqWords(b *testing.B) {
	word := words[100]

	// Run the test
	for n := 0; n < b.N; n++ {
		var count int
		for _, w := range words {
			if strings.EqualFold(word, w) {
				count += 1
			}
		}
		if count != 1 {
			b.Error("Failed to find the match")
		}
	}
}

func BenchmarkGlobWords(b *testing.B) {
	benchmarkGlob(words, wordsCount, b)
}
func BenchmarkGreedyRegexWords(b *testing.B) {
	benchmarkRegexWords(words, greedyRegexPatterns, wordsCount, b)
}
func BenchmarkNonGreedyRegexWords(b *testing.B) {
	benchmarkRegexWords(words, nonGreedyRegexPatterns, wordsCount, b)
}

func BenchmarkGlobLogLines(b *testing.B) {
	benchmarkGlob(logLines, logLinesCount, b)
}
func BenchmarkGreedyRegexLogLines(b *testing.B) {
	benchmarkRegexWords(logLines, greedyRegexPatterns, logLinesCount, b)
}
func BenchmarkNonGreedyRegexLogLines(b *testing.B) {
	benchmarkRegexWords(logLines, nonGreedyRegexPatterns, logLinesCount, b)
}

func benchmarkGlob(lines []string, count int, b *testing.B) {

	// Run the test
	for n := 0; n < b.N; n++ {
		globTest(globPatterns, lines, count, b)
	}
}

func benchmarkRegexWords(lines []string, patterns []string, count int, b *testing.B) {

	// Run the test
	for n := 0; n < b.N; n++ {
		regexTest(greedyRegexPatterns, lines, count, b)
	}
}

func readBenchmarkFile(name string) []string {
	var lines []string
	file, err := os.Open("benchmark/" + name)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}

func globTest(globPatterns []string, lines []string, assertCount int, b *testing.B) {
	count := 0
	for _, pattern := range globPatterns {
		m, err := Compile(pattern, WithWildcardChar('*'), WithoutMatchOne(), WithCaseInsensitive(true))
		if err != nil {
			panic(err)
		}
		for _, line := range lines {
			if m.Matches(line) {
				count += 1
			}
		}
	}
	// To keep us honest
	if count != assertCount {
		b.Errorf("Count was not equal %d != %d", count, assertCount)
	}
}

func regexTest(regexPatterns []string, lines []string, assertCount int, b *testing.B) {
	count := 0
	for _, pattern := range regexPatterns {
		r, _ := regexp.Compile(pattern)
		for _, line := range lines {
			if r.MatchString(line) {
				count += 1
			}
		}
	}
	// To keep us honest
	if count != assertCount {
		b.Errorf("Count was not equal %d != %d", count, assertCount)
	}
}
