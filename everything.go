package ktanbu

import (
	"fmt"
)

// ComplexWires - Runs the complex wires module
func ComplexWires() {
	order := []bool{}
	cWires := make(map[string]string)
	cWires = map[string]string{
		"woo": "c",
		"wol": "d",
		"wso": "c",
		"wsl": "b",
		"roo": "s",
		"rol": "b",
		"rso": "c",
		"rsl": "b",
		"boo": "s",
		"bol": "p",
		"bso": "d",
		"bsl": "p",
		"poo": "s",
		"pol": "s",
		"pso": "p",
		"psl": "d",
	}

	fmt.Printf("Please enter the pattern of wires: ")
	var combinations []string
	combination := inputString()
	for i := 0; i < len(combination); i = i + 3 {
		combinations = append(combinations, combination[i:i+3])
	}
	for _, i := range combinations {
		switch cWires[i] {
		case "c":
			order = append(order, true)
		case "d":
			order = append(order, false)
		case "b":
			recieveBatteries()
			if batteries >= 2 {
				order = append(order, true)
			} else {
				order = append(order, false)
			}
		case "s":
			recieveSerial()
			if serialEven {
				order = append(order, true)
			} else {
				order = append(order, false)
			}
		case "p":
			recieveParallel()
			if parallelPort {
				order = append(order, true)
			} else {
				order = append(order, false)
			}
		}
	}
	fmt.Printf("Cut the wire(s) in ")
	for index, cut := range order {
		if cut {
			fmt.Printf("%s, ", Position[index])
		}
	}
	fmt.Printf("positions.\n")
}
