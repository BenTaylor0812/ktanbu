package ktanbu

import "fmt"

func obtainComb() string {
	var wires string
	fmt.Printf("Enter the combinations of wires (reading down): ")
	for {
		cVal := true
		wires = inputString()
		if len(wires) > 6 {
			fmt.Printf("That is too long, type again: ")
			continue
		}
		if len(wires) < 3 {
			fmt.Printf("That is too short, type again: ")
			continue
		}
		for i, val := range wires {
			letter := string(val)
			switch letter {
			case "k", "w", "b", "r", "y":
				continue
			default:
				if cVal {
					fmt.Printf("The value %v in the %s position is incorrect. type again: ", letter, Position[i])
					cVal = false
				}
			}
		}
		if cVal {
			break
		}
	}
	return wires
}

// Wires - Runs the wires module
func Wires() {
	wires := obtainComb()
	wireLength := len(wires)
	lastWire := string(wires[wireLength-1])
	red := 0
	blue := 0
	yellow := 0
	black := 0
	white := 0
	lastBlue := 0
	lastRed := 0
	for i, val := range wires {
		if string(val) == "b" {
			blue++
			lastBlue = i
		} else if string(val) == "r" {
			red++
			lastRed = i
		} else if string(val) == "y" {
			yellow++
		} else if string(val) == "k" {
			black++
		} else if string(val) == "w" {
			white++
		}

	}
	switch wireLength {
	case 3:
		if red == 0 {
			fmt.Println(cutSecond)
		} else if string(wires[wireLength-1]) == "w" {
			fmt.Println(cutLast)
		} else if blue > 1 {
			fmt.Printf("Cut last blue wire. (The %s wire.)\n", Position[lastBlue])
		} else {
			fmt.Println(cutLast)
		}
	case 4:
		recieveSerial()
		if serialOdd && red > 1 {
			fmt.Printf("Cut the last red wire. (The %s wire.)\n", Position[lastRed])
		} else if lastWire == "y" && red == 0 {
			fmt.Println(cutFirst)
		} else if blue == 1 {
			fmt.Println(cutFirst)
		} else if yellow > 1 {
			fmt.Println(cutLast)
		} else {
			fmt.Println(cutSecond)
		}
	case 5:
		recieveSerial()
		if lastWire == "k" && serialOdd {
			fmt.Print(cutFourth)
		} else if red == 1 && yellow > 1 {
			fmt.Println(cutFirst)
		} else if black == 0 {
			fmt.Println(cutSecond)
		} else {
			fmt.Println(cutFirst)
		}
	case 6:
		if yellow == 0 && serialOdd {
			fmt.Println(cutThird)
		} else if yellow == 1 && white > 1 {
			fmt.Println(cutFourth)
		} else if red == 0 {
			fmt.Println(cutLast)
		} else {
			fmt.Print(cutFourth)
		}
	}
}
