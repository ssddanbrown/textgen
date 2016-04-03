package main

import (
	"math"
	"math/rand"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type textGenerator struct {
	wordList           []string
	wordCount          int
	wordsPerLine       int
	linesPerParagraph  int
	numberOfParagraphs int
	variation          float64
}

func (w *textGenerator) addVariation(inputVal int) int {
	if w.variation == 0 {
		return inputVal
	}

	inputValF := float64(inputVal)

	minVal := math.Max(inputValF*w.variation, 1)
	maxVal := inputValF * (1 + w.variation)

	valDiff := math.Max(maxVal-minVal, 1)

	return rand.Intn(int(valDiff)) + int(minVal)
}

func (w *textGenerator) genLine() string {
	final := ""
	wordsPerLine := w.addVariation(w.wordsPerLine)

	for i := 0; i < wordsPerLine; i++ {
		randomPosition := rand.Intn(w.wordCount)
		var addWord string

		if i == 0 {
			addWord = strings.Title(w.wordList[randomPosition])
		} else {
			addWord = strings.ToLower(w.wordList[randomPosition])
		}

		final += addWord + " "
	}

	return strings.TrimRight(final, " ") + "."
}

func (w *textGenerator) genParagraph() string {
	final := ""
	linesPerParagraph := w.addVariation(w.linesPerParagraph)
	for i := 0; i < linesPerParagraph; i++ {
		final += w.genLine() + " "
	}
	return final
}

func (w *textGenerator) genParagraphs() string {
	final := ""
	numberOfParagraphs := w.addVariation(w.numberOfParagraphs)
	for i := 0; i < numberOfParagraphs; i++ {
		final += w.genParagraph() + "\n\n"
	}
	return strings.TrimRight(final, "\n")
}

func newTextGenerator() *textGenerator {
	rand.Seed(time.Now().UTC().UnixNano())
	textGenerator := textGenerator{}
	// Load in words
	wordList := wordListString
	// Split words ready to use
	textGenerator.wordList = strings.Split(string(wordList), "\n")
	textGenerator.wordCount = len(textGenerator.wordList)
	// Return our new text generator
	return &textGenerator
}
