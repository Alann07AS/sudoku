package sudoFonction

func IsInSquare(x, y, value int, table [9][9]int) bool {
	const X, Y = 0, 1
	var squareRange = []int{0, 0}
	switch x {
	case 0, 1, 2:
		squareRange[X] = 2
	case 3, 4, 5:
		squareRange[X] = 5
	case 6, 7, 8:
		squareRange[X] = 8
	}
	switch y {
	case 0, 1, 2:
		squareRange[Y] = 2
	case 3, 4, 5:
		squareRange[Y] = 5
	case 6, 7, 8:
		squareRange[Y] = 8
	}
	for indexY := squareRange[Y] - 2; indexY <= squareRange[Y]; indexY++ {
		for indexX := squareRange[X] - 2; indexX <= squareRange[X]; indexX++ {
			if table[indexY][indexX] == value {
				return true
			}
		}
	}
	return false
}
