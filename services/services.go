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

func FindValidSequence(matrix [][]bool) int {
	sequencesFounded := 0
	for indexRow, row := range matrix {
		for indexColumn := range row {
			if CheckForDiagonalSequence(matrix, indexRow, indexColumn) {
				sequencesFounded++
				continue
			}
			if CheckForLineSequence(matrix, indexRow, indexColumn) {
				sequencesFounded++
				continue
			}
			if CheckForColumnSequence(matrix, indexRow, indexColumn) {
				sequencesFounded++
				continue
			}
		}
	}

	return sequencesFounded
}

func CheckForDiagonalSequence(matrix [][]bool, actualRow, actualColumn int) bool {
	countLetter := 0
	maxRows, maxColumns := GetMatrixSize(matrix)

	row := actualRow
	column := actualColumn
	for row < maxRows && column < maxColumns {
		if matrix[row][column] {
			countLetter++
		} else {
			break
		}
		if countLetter == 4 {
			return true
		}
		row++
		column++
	}

	countLetter = 0
	row = actualRow
	column = actualColumn
	for row < maxRows && column >= 0 {
		if matrix[row][column] {
			countLetter++
		} else {
			break
		}
		if countLetter == 4 {
			return true
		}
		row++
		column--
	}
	return false
}

func CheckForLineSequence(matrix [][]bool, actualRow, actualColumn int) bool {
	countLetter := 0
	_, maxColumns := GetMatrixSize(matrix)
	for column := actualColumn; column < maxColumns; column++ {
		if matrix[actualRow][column] {
			countLetter++
		} else {
			break
		}
		if countLetter == 4 {
			return true
		}
	}
	return false
}

func CheckForColumnSequence(matrix [][]bool, actualRow, actualColumn int) bool {
	countLetter := 0
	maxRows, _ := GetMatrixSize(matrix)
	for row := actualRow; row < maxRows; row++ {
		if matrix[row][actualColumn] {
			countLetter++
		} else {
			break
		}
		if countLetter == 4 {
			return true
		}
	}
	return false
}

func GetMatrixSize(matrix [][]bool) (int, int) {
	maxRows := len(matrix)
	var maxColumns int
	if maxRows > 0 {
		maxColumns = len(matrix[0])
	} else {
		maxColumns = 0
	}
	return maxRows, maxColumns
}
