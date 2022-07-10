package sudoFonction

func IsCorectArgument(argsTable []string) (bool, int, int) {
	correctArg := true
	caseLost := 0
	if len(argsTable) != 10 {
		correctArg = false
		return correctArg, len(argsTable) - 10, caseLost
	}
	for i := 1; i <= 9; i++ {
		if len(argsTable[i]) != 9 {
			correctArg = false
			if len(argsTable[i])-9 < 0 {
				caseLost += (len(argsTable[i]) - 9) * -1
			} else {
				caseLost += len(argsTable[i]) - 9
			}
		}
	}
	return correctArg, len(argsTable) - 10, caseLost
}
