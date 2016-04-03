package main

import (
	"flag"
	"fmt"
)

func main() {
	// Options
	optWordCount := flag.Int("w", 10, "Words per line")
	optLineCount := flag.Int("l", 12, "Lines per paragraph")
	optParagraphCount := flag.Int("p", 4, "Number of paragraphs")
	optRandomness := flag.Float64("r", 0.5, "Randomness/Variations in lengths")
	flag.Parse()

	// Create return string
	generator := newTextGenerator()
	generator.wordsPerLine = *optWordCount
	generator.linesPerParagraph = *optLineCount
	generator.numberOfParagraphs = *optParagraphCount
	generator.variation = *optRandomness

	// TODO - prevent this blocking up
	// Write out to command line as creating
	//
	// Add testing and benchmarks beforehand
	line := generator.genParagraphs()
	fmt.Println(line)
}
