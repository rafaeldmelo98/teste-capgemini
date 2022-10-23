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
	sequencesFound := 0
	var elementsFoundCoordinates [][]int
	for indexRow, row := range matrix {
		for indexColumn := range row {
			elementvalidated := CheckIfElementAlreadyValidated(elementsFoundCoordinates, indexRow, indexColumn)
			if elementvalidated {
				continue
			}
			foundDiagonalSequence, diagonalDirection := CheckForDiagonalSequence(matrix, indexRow, indexColumn)
			if foundDiagonalSequence {
				elementsFoundCoordinates = append(elementsFoundCoordinates,
					[]int{indexRow, indexColumn},
					[]int{indexRow + 1, indexColumn + (diagonalDirection * 1)},
					[]int{indexRow + 2, indexColumn + (diagonalDirection * 2)},
					[]int{indexRow + 3, indexColumn + (diagonalDirection * 3)})
				sequencesFound++
				continue
			}
			if CheckForLineSequence(matrix, indexRow, indexColumn) {
				elementsFoundCoordinates = append(elementsFoundCoordinates,
					[]int{indexRow, indexColumn}, []int{indexRow, indexColumn + 1},
					[]int{indexRow, indexColumn + 2}, []int{indexRow, indexColumn + 3})
				sequencesFound++
				continue
			}
			if CheckForColumnSequence(matrix, indexRow, indexColumn) {
				elementsFoundCoordinates = append(elementsFoundCoordinates,
					[]int{indexRow, indexColumn}, []int{indexRow + 1, indexColumn},
					[]int{indexRow + 2, indexColumn}, []int{indexRow + 3, indexColumn})
				sequencesFound++
				continue
			}
		}
	}

	return sequencesFound
}

func CheckForDiagonalSequence(matrix [][]bool, actualRow, actualColumn int) (bool, int) {
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
			return true, +1
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
			return true, -1
		}
		row++
		column--
	}
	return false, 0
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

func CheckIfElementAlreadyValidated(elementsFound [][]int, actualRow, actualColumn int) bool {
	for _, coordinates := range elementsFound {
		if coordinates[0] == actualRow && coordinates[1] == actualColumn {
			return true
		}
	}
	return false
}
