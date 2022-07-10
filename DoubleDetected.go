package sudoFonction

import "fmt"

func DoubleDectected(intTable [9][9]int) bool {
	//double dans les lignes
	for ligne := 0; ligne <= 8; ligne++ {
		table := sortLigne(ligne, intTable)
		fmt.Println("ligne ", table)
		for i := 0; i <= len(table)-2; i++ {
			if table[i] == table[i+1] {
				return true
			}
		}
	}

	//double colone
	for ligne := 0; ligne <= 8; ligne++ {
		table := sortRow(ligne, intTable)
		fmt.Println("row ", table)
		for i := 0; i <= len(table)-2; i++ {
			if table[i] == table[i+1] {
				return true
			}
		}
	}

	//double dans caré
	for ligne := 0; ligne <= 8; ligne++ {
		table := sortSquare(ligne, intTable)
		fmt.Println("ligne ", table)
		for i := 0; i <= len(table)-2; i++ {
			if table[i] == table[i+1] {
				return true
			}
		}
	}
	return false
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
			fmt.Print(indexX, "__", nbL, "   ")
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
