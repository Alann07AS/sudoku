package sudoFonction

func IsInHorizontal(y, value int, tab [9][9]int) bool {
	for i := 0; i <= 8; i++ {
		if value == tab[y][i] {
			return true
		}
	}
	return false
}
