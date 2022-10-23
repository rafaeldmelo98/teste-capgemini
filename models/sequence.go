package models

type Sequence struct {
	SequenceList []string `json:"letters"`
}

type SequenceTable struct {
	QuantityValidSequence   int     `json:"count_valid"`
	QuantityInvalidSequence int     `json:"count_invalid"`
	RateValidSequence       float64 `json:"ratio"`
}
