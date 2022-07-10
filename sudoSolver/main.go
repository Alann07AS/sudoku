package main

import (
	"fmt"
	"os"
	"sudoFonction"
)

func main() {
	argumentTable := os.Args
	if !sudoFonction.IsCorectArgument(argumentTable) {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("!!!!!!!BOARD NOT VALID!!!!!!!")
		fmt.Println("!!!!!!!ARGURMENT NUMBER!!!!!!")
		fmt.Println("!!!!!!!!!!!!FALSE!!!!!!!!!!!!")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		return
	}
	var grilleTable = sudoFonction.PutInIntTable(argumentTable)
	if sudoFonction.DoubleDectected(grilleTable) {
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		fmt.Println("!!!!!!!BOARD NOT VALID!!!!!!!")
		fmt.Println("!!!!!!!!DOUBLE NUMBER!!!!!!!!")
		fmt.Println("!!!!!!!!!!!DETECTED!!!!!!!!!!!")
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
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
					fmt.Println("Position", indexX, " ", indexY, "possibleValue", possibleValue)
					if len(possibleValue) == 1 {
						grilleTable[indexY][indexX] = possibleValue[0]
					}
					fmt.Println(allPossibleValueInLigne)

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
			fmt.Println(sudoFonction.DeductiveSolve(allPossibleValueInLigne))
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
