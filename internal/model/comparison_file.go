package model

type ComparisonFile struct {
	Label string
	Data  []byte
}

func NewComparisonFile(label string, data []byte) *ComparisonFile {
	return &ComparisonFile{Label: label, Data: data}
}
