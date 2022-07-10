package sudoFonction

func DeductiveSolve(table []int) []int {
	if len(table) == 0 {
		return table
	}
	isModif := true
	for isModif {
		isModif = false
		for i := 0; i <= len(table)-2; i++ {
			if table[i] < table[i+1] {
				isModif = true
				a := table[i]
				table[i] = table[i+1]
				table[i+1] = a
			}
		}
	}
	
	last := table[0]
	for i := 0; i <= len(table)-2; i++ {
		if last == table[i+1] {
			last = table[i+1]
			table[i] = 0
			table[i+1] = 0
		} else {
			last = table[i+1]
		}
	}
	var newTable []int
	for _, nb := range table {
		if nb != 0 {
			newTable = append(newTable, nb)
		}
	}

	return newTable
}
