package sudoFonction

func IsInVertical(x, value int, tab [9][9]int) bool {
	for i := 0; i <= 8; i++ {
		if value == tab[x][i] {
			return true
		}
	}
	return false
}
