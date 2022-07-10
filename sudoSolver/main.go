package main

import (
	"fmt"
	"os"
	"sudoFonction"
)

func main() {
	argumentTable := os.Args
	fmt.Println()
	isCorectArg, argLost, caseLost := sudoFonction.IsCorectArgument(argumentTable)
	if !isCorectArg {
		fmt.Println()
		fmt.Println("They are argument mistake:")
		fmt.Println()
		if argLost != 0 {
			fmt.Println("They are", argLost, "arguments")
		}
		if caseLost != 0 {
			fmt.Println("They are", caseLost, "case of dif")
		}
		fmt.Println()
		return
	}
	var grilleTable = sudoFonction.PutInIntTable(argumentTable)
	isDoublon, dL, dR, dS := sudoFonction.DoubleDectected(grilleTable)
	if isDoublon {
		fmt.Println()
		fmt.Println("Errors there are", dL+dR+dS, "double :")
		fmt.Println()
		fmt.Println("-One or more double in", dL, "ligne.")
		fmt.Println("-One or more double in", dR, "row.")
		fmt.Println("-One or more double in", dS, "square.")
		fmt.Println()
		return
	}

	i := 0
	for i != 50 {
		fmt.Println("____________________________________________________________________")
		for indexY := 0; indexY <= 8; indexY++ {
			var allPossibleValueInLigne []int
			for indexX := 0; indexX <= 8; indexX++ {
				// deductValue := sudoFonction.DeductiveSolve(allPossibleValueInLigne)
				if grilleTable[indexY][indexX] == 0 {
					var possibleValue []int
					for value := 1; value <= 9; value++ {
						if isPossible(indexX, indexY, value, grilleTable) {
							// fmt.Println("is POssible")
							possibleValue = append(possibleValue, value)
							allPossibleValueInLigne = append(allPossibleValueInLigne, value)
						}
					}
					if len(possibleValue) == 1 {
						grilleTable[indexY][indexX] = possibleValue[0]
					}

				}
				// if len(sudoFonction.DeductiveSolve(allPossibleValueInLigne)) != 0 {
				// 	indexX = 0
				// }

			}
			tableDeduct := sudoFonction.DeductiveSolve(allPossibleValueInLigne)
			if len(tableDeduct) != 0 {
				for i := 0; i <= len(tableDeduct)-1; i++ {
					for indexX := 0; indexX <= 8; indexX++ {
						if grilleTable[indexY][indexX] == 0 {
							if isPossible(indexX, indexY, tableDeduct[i], grilleTable) {
								grilleTable[indexY][indexX] = tableDeduct[i]
							}
						}
					}
				}
			}
		}
		i++
	}
	displayTest(grilleTable)
}

func displayTest(grilleTable [9][9]int) {
	for _, row := range grilleTable {
		fmt.Println(row)
	}
	fmt.Println()
}

func isPossible(x, y, value int, table [9][9]int) bool {
	if sudoFonction.IsInHorizontal(y, value, table) ||
		sudoFonction.IsInVertical(x, value, table) ||
		sudoFonction.IsInSquare(x, y, value, table) {
		return false
	} else {
		return true
	}
}
