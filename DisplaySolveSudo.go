package sudoFonction

import "fmt"

func DisplaySolveSudo(sudo [9][9]int) {
	verticalPos := -1
	horizontalPos := -1
	for i := 0; i <= 34; i++ {
		fmt.Print("—")
	}
	fmt.Println()
	for y := 0; y <= 8; y++ {
		fmt.Print("|")
		fmt.Print("|")
		verticalPos = -1
		for x := 0; x <= 8; x++ {
			fmt.Print(" ", sudo[y][x], " ")
			if verticalPos+3 == x {
				fmt.Print("|")
				fmt.Print("|")
				verticalPos = x
			}
		}
		fmt.Println()
		verticalPos = -1
		fmt.Print("|")
		fmt.Print("|")
		for void := 0; void <= 8; void++ {
			fmt.Print("   ")
			if verticalPos+3 == void {
				fmt.Print("|")
				fmt.Print("|")
				verticalPos = void
			}
		}
		verticalPos = -1
		if horizontalPos+3 == y {
			fmt.Println()
			horizontalPos = y
			for i := 0; i <= 34; i++ {
				fmt.Print("—")
			}
		}
		fmt.Println()
	}
}
