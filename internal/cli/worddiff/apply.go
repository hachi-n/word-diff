package worddiff

import (
	"github.com/hachi-n/word-diff/internal/loader"
	"github.com/hachi-n/word-diff/internal/model"
)

func Apply(fileName1, fileName2 string) error {
	comparator := model.NewComparator(
		model.NewComparisonFile(fileName1, loader.LoadWithExitOnFail(fileName1)),
		model.NewComparisonFile(fileName2, loader.LoadWithExitOnFail(fileName2)),
	)
	return comparator.PrintWordDiff()
}
