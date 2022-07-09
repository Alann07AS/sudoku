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
		fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
		return
	}
	var grilleTable = sudoFonction.PutInIntTable(argumentTable)
	displayTest(grilleTable)
	fmt.Println(sudoFonction.IsInSquare(8, 8, 1, grilleTable))
}

func displayTest(grilleTable [9][9]int) {
	for _, row := range grilleTable {
		fmt.Println(row)
	}
	fmt.Println()
}
