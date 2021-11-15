package model

import (
	"fmt"
	"strings"
)

type Comparator struct {
	ComparisonFile1 *ComparisonFile
	ComparisonFile2 *ComparisonFile
}

func NewComparator(comparisonFile1 *ComparisonFile, comparisonFile2 *ComparisonFile) *Comparator {
	return &Comparator{ComparisonFile1: comparisonFile1, ComparisonFile2: comparisonFile2}
}

func (c *Comparator) PrintWordDiff() error {
	leftPadding := 0
	rightPadding := 0
	if c.ComparisonFile1.Label < c.ComparisonFile2.Label {
		leftPadding = len(c.ComparisonFile2.Label) - len(c.ComparisonFile1.Label)
	} else {
		rightPadding = len(c.ComparisonFile1.Label) - len(c.ComparisonFile2.Label)
	}

	checkStrictsLength := 10
	displayWordLength := 100

	differencePointIndex := 0
	tmpData := make([]byte, len(c.ComparisonFile2.Data))
	copy(tmpData, c.ComparisonFile2.Data)

	for i, v := range c.ComparisonFile1.Data {
		if v == tmpData[i] {
			continue
		}

		// difference pattern.
		differencePointIndex = i

		fmt.Printf("=========== word-diff::: %s: %d ==========\n", c.ComparisonFile1.Label, differencePointIndex)
		fmt.Printf("index:%d, value:%s \n", differencePointIndex, string(v))
		fmt.Printf("%s%s: %s \n", c.ComparisonFile1.Label, strings.Repeat(" ", leftPadding), string(c.ComparisonFile1.Data[i:i+displayWordLength]))
		fmt.Printf("%s%s: %s \n", c.ComparisonFile2.Label, strings.Repeat(" ", rightPadding), string(tmpData[i:i+displayWordLength]))
		fmt.Printf("=========== word-diff::: %s: %d ==========\n", c.ComparisonFile1.Label, differencePointIndex)

		incr := c.checkStrictsWords(c.ComparisonFile1.Data[i:i+checkStrictsLength], tmpData[i:])
		// offset
		tmpData = append(tmpData[:differencePointIndex], tmpData[differencePointIndex+incr:]...)
	}

	return nil
}

func (c *Comparator) checkStrictsWords(testWords []byte, words []byte) int {
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
