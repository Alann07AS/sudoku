package sudoFonction

func DoubleDectected(intTable [9][9]int) (bool, int, int, int) {
	doublonLigne := 0
	doublonRow := 0
	doublonSquare := 0
	isDoublon := false

	//double dans les lignes
	for ligne := 0; ligne <= 8; ligne++ {
		table := sortLigne(ligne, intTable)
		for i := 0; i <= len(table)-2; i++ {
			if table[i] == table[i+1] {
				isDoublon = true
				doublonLigne++
			}
		}
	}

	//double colone
	for ligne := 0; ligne <= 8; ligne++ {
		table := sortRow(ligne, intTable)
		for i := 0; i <= len(table)-2; i++ {
			if table[i] == table[i+1] {
				isDoublon = true
				doublonRow++
			}
		}
	}

	//double dans caré
	for ligne := 0; ligne <= 8; ligne++ {
		table := sortSquare(ligne, intTable)
		for i := 0; i <= len(table)-2; i++ {
			if table[i] == table[i+1] {
				isDoublon = true
				doublonSquare++
			}
		}
	}
	return isDoublon, doublonLigne, doublonRow, doublonSquare
}

func isNbrInt(nb int) bool {
	if nb >= 1 && nb <= 9 {
		return true
	} else {
		return false
	}
}

func sortLigne(nbL int, intTable [9][9]int) []int {
	var table []int
	for indexX := 0; indexX <= 8; indexX++ {
		if isNbrInt(intTable[nbL][indexX]) {
			table = append(table, intTable[nbL][indexX])
		}
	}
	return sort(table)
}

func sortRow(nbR int, intTable [9][9]int) []int {
	var table []int
	for indexY := 0; indexY <= 8; indexY++ {
		if isNbrInt(intTable[indexY][nbR]) {
			table = append(table, intTable[indexY][nbR])
		}
	}
	return sort(table)
}

func sortSquare(nbSqr int, intTable [9][9]int) []int {
	const X, Y = 0, 1
	var squareRange = []int{0, 0}
	switch nbSqr {
	case 0, 3, 6:
		squareRange[X] = 2
	case 1, 4, 7:
		squareRange[X] = 5
	case 2, 5, 8:
		squareRange[X] = 8
	}
	switch nbSqr {
	case 0, 1, 2:
		squareRange[Y] = 2
	case 3, 4, 5:
		squareRange[Y] = 5
	case 6, 7, 8:
		squareRange[Y] = 8
	}
	var table []int
	for indexY := squareRange[Y] - 2; indexY <= squareRange[Y]; indexY++ {
		for indexX := squareRange[X] - 2; indexX <= squareRange[X]; indexX++ {
			if isNbrInt(intTable[indexY][indexX]) {
				table = append(table, intTable[indexY][indexX])
			}
		}
	}
	return sort(table)
}

func sort(table []int) []int {
	isModified := true
	for isModified {
		isModified = false
		for i := 0; i <= len(table)-2; i++ {
			if table[i] < table[i+1] {
				a := table[i]
				table[i] = table[i+1]
				table[i+1] = a
				isModified = true
			}
		}
	}
	return table
}
