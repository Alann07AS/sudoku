package sudoFonction

//Fonction de transeformation de tableau de string en tableau de int

func PutInIntTable(argsTable []string) [9][9]int {
	var intTable [9][9]int
	for i := 1; i <= 9; i++ {
		charTable := []rune(argsTable[i])
		for i2, char := range charTable {
			if isNbr(char) {
				intTable[i-1][i2] = int(char) - 48
			} else {
				intTable[i-1][i2] = 0
			}
		}
	}

	return intTable
}

func isNbr(char rune) bool {
	if char >= '1' && char <= '9' {
		return true
	} else {
		return false
	}
}

func doubleDectected(intTable [9][9]int) bool {
	for i := 0; i <= 8; i++ {
		for i2 := 0; i2 <= 8; i2++ {
			if intTable[i][i2] != 0 {
				if IsInSquare(i, i2, intTable[i][i2], intTable) {

				}
			}
		}
	}
	return true
}
