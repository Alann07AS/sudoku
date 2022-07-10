package sudoFonction

func IsSolve(table [9][9]int) bool {
	for y := 0; y <= 8; y++ {
		for x := 0; x <= 8; x++ {
			if table[y][x] == 0 {
				return false
			}
		}
	}
	return true
}
