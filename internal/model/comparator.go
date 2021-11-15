package model

import (
	"fmt"
	"strings"
)

type Comparator struct {
	ComparisonFile1 *ComparisonFile
	ComparisonFile2 *ComparisonFile
}

const (
	// a value that is neither too long nor too short
	checkWordLength = 10

	// Length of characters in comparison result
	displayWordLength = 100
)

func NewComparator(comparisonFile1 *ComparisonFile, comparisonFile2 *ComparisonFile) *Comparator {
	return &Comparator{ComparisonFile1: comparisonFile1, ComparisonFile2: comparisonFile2}
}

func (c *Comparator) PrintWordDiff() error {
	leftPadding := 0
	rightPadding := 0

	file1LabelLength := len(c.ComparisonFile1.Label)
	file2LabelLength := len(c.ComparisonFile2.Label)
	if file1LabelLength < file2LabelLength {
		leftPadding = file2LabelLength - file1LabelLength
	} else {
		rightPadding = file1LabelLength - file2LabelLength
	}

	differencePointIndex := 0
	tmpData := make([]byte, len(c.ComparisonFile2.Data))
	copy(tmpData, c.ComparisonFile2.Data)

	count := 1

	for i, v := range c.ComparisonFile1.Data {
		if v == tmpData[i] {
			continue
		}

		// difference pattern.
		differencePointIndex = i

		fmt.Printf("=========== %d. word-diff::: %s: %d ==========\n", count, c.ComparisonFile1.Label, differencePointIndex)
		fmt.Printf("index:%d, value:%s \n", differencePointIndex, string(v))
		fmt.Printf("%s%s: %s \n", c.ComparisonFile1.Label, strings.Repeat(" ", leftPadding), string(c.ComparisonFile1.Data[i:i+displayWordLength]))
		fmt.Printf("%s%s: %s \n", c.ComparisonFile2.Label, strings.Repeat(" ", rightPadding), string(tmpData[i:i+displayWordLength]))
		fmt.Printf("=========== %d. word-diff::: %s: %d ==========\n", count, c.ComparisonFile1.Label, differencePointIndex)

		incr := c.checkSameWords(c.ComparisonFile1.Data[i:i+checkWordLength], tmpData[i:])
		// offset
		tmpData = append(tmpData[:differencePointIndex], tmpData[differencePointIndex+incr:]...)
		count++
	}

	return nil
}

func (c *Comparator) checkSameWords(testWords []byte, words []byte) int {
	tmpIndex := 0
	testWordsIndex := 0
	firstSamePointIndex := 0

	for tmpIndex < (len(words)-1) && len(testWords) != testWordsIndex {
		if testWords[testWordsIndex] == words[tmpIndex] {
			if firstSamePointIndex == 0 {
				firstSamePointIndex = tmpIndex
			}
			testWordsIndex++
		} else {
			testWordsIndex = 0
			firstSamePointIndex = 0
		}
		tmpIndex++
	}
	return firstSamePointIndex
}
