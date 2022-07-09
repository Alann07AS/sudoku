package sudoFonction

func IsCorectArgument(argsTable []string) bool {
	if len(argsTable) != 10 {
		return false
	}
	for i := 1; i <= 9; i++ {
		if len(argsTable[i]) != 9 {
			return false
		}
	}
	return true
}
