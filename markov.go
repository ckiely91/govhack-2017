package main

import (
	"math/rand"
	"strings"
)

type runeCounts struct {
	countMap   map[rune]int
	totalCount int
}

func (counts *runeCounts) addRune(c rune) {
	if _, ok := counts.countMap[c]; ok {
		counts.countMap[c]++
	} else {
		counts.countMap[c] = 1
	}
	counts.totalCount++
}

func (counts *runeCounts) getRandomRune() rune {
	// Use count weightings to determine
	var probSum float64 = 0
	randProb := rand.Float64()

	var chosenChar rune
	for chosenChar, count := range counts.countMap {
		thisProbRange := float64(count) / float64(counts.totalCount)
		if probSum < randProb && randProb < probSum+thisProbRange {
			return chosenChar
		}
		probSum += thisProbRange
	}

	return chosenChar
}

type Markov struct {
	chain map[string]runeCounts
	n     int
}

func NewMarkov(n int) *Markov {
	return &Markov{
		chain: map[string]runeCounts{},
		n:     n,
	}
}

func (m *Markov) ParseWord(word string) {
	runes := []rune(word)
	if m.n > len(runes) {
		return
	}

	end := len(runes) - m.n
	for i := 0; i < end; i++ {
		key := string(runes[i : i+m.n])
		val := runes[i+m.n]
		runeCount, ok := m.chain[key]
		if ok {
			runeCount.addRune(val)
		} else {
			m.chain[key] = runeCounts{
				countMap:   map[rune]int{val: 1},
				totalCount: 1,
			}
		}
	}
}

func (m *Markov) GenerateWord(length int) string {
	letters := []rune(m.getRandomPrefix())

	for i := m.n; i <= length; i++ {
		lastLetters := letters[i-m.n : i]
		// fmt.Printf("Looking for key %v\n", string(lastLetters))
		if runes, ok := m.chain[string(lastLetters)]; ok {
			letters = append(letters, runes.getRandomRune())
		} else {
			return string(letters)
		}
	}

	return string(letters)
}

func (m *Markov) GenerateBusinessName() string {
	numWords := rand.Intn(3) + 1
	words := []string{}

	for i := 0; i < numWords; i++ {
		wordLength := rand.Intn(7) + 4
		word := m.GenerateWord(wordLength)
		word = strings.Title(strings.ToLower(word))
		words = append(words, word)
	}

	return strings.Join(words, " ")
}

func (m *Markov) getRandomPrefix() string {
	// Go randomises the order of maps every time
	for key, _ := range m.chain {
		return key
	}

	return "ab"
}

func getRandomLetter(slice []rune) rune {
	return slice[rand.Intn(len(slice))]
}
