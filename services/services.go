package services

func MapSequence(letter string, sequenceList []string) [][]bool {
	var matrix [][]bool
	for _, sequence := range sequenceList {
		var row []bool
		for _, letterSequence := range sequence {
			if string(letterSequence) == letter {
				row = append(row, true)
			} else {
				row = append(row, false)
			}
		}
		matrix = append(matrix, row)
	}

	return matrix
}
