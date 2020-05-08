package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	wordLength int = 5
)

const (
	cutFirst  string = "Cut the first wire."
	cutSecond string = "Cut the second wire."
	cutThird  string = "Cut the third wire."
	cutFourth string = "Cut the fourth wire."
	cutFifth  string = "Cut the fifth wire."
	cutLast   string = "Cut the last wire."
)

type gameFunction = func()

var position []string
var reader *bufio.Reader
var serial string
var serialOdd bool
var serialEven bool
var batteries int
var batteryLock bool
var parallelPort bool
var parallelLock bool

func main() {
	reader = bufio.NewReader(os.Stdin)
	position = []string{"first", "second", "third", "fourth", "fifth"}
	complexWires()
	// for gameOver := false; !gameOver; {
	// 	fmt.Printf("Enter the game you are playing: ")
	// 	game, _ := reader.ReadString('\n')
	// 	game = strings.Replace(game, "\r\n", "", -1)
	// 	switch game {
	// 	case "word", "words":
	// 		wordGame()
	// 	case "wires", "wire":
	// 		wires()
	// 	case "cwires", "cwire", "complexwires", "complex wires":
	// 		complexWires()
	// 	case "win", "won":
	// 		fmt.Println("Well done!")
	// 	case "loss", "lost", "lose":
	// 		fmt.Println("Ah sorry, maybe next time!")
	// 	case "leave":
	// 		fmt.Println("Hope you did well!")
	// 		gameOver = true
	// 		break
	// 	default:
	// 		fmt.Println("Sorry, didn't catch that?")
	// 	}
	// }
}

func yesOrNo() string {
	var response string
	for {
		response = inputString()
		switch response {
		case "yes", "y":
			return "yes"
		case "no", "n":
			fmt.Println("Okay then.")
			return "no"
		default:
			fmt.Printf("Please enter yes or no: ")
			continue
		}
	}
}

func inputString() string {
	output, _ := reader.ReadString('\n')
	return strings.Replace(output, "\r\n", "", -1)
}

func inputInteger() int {
	var output string
	for {
		output = inputString()
		rVal, err := strconv.Atoi(output)
		if err != nil {
			fmt.Printf("Please only enter an integer: ")
		} else {
			return rVal
		}
	}

}

func recieveSerial() {
	for serial == "" {
		fmt.Printf("Enter the serial number: ")
		serial = inputString()
		serialOdd = int(serial[len(serial)-1])%2 == 1
		serialEven = !serialOdd
	}
}

func recieveParallel() {
	for !parallelLock {
		fmt.Printf("Is there a parallel port? ")
		response := yesOrNo()
		if response == "yes" {
			parallelPort = true
		} else {
			parallelPort = false
		}
		parallelLock = true
	}
}

func recieveBatteries() {
	for !batteryLock {
		fmt.Printf("How many batteries are there? ")
		batteries = inputInteger()
		batteryLock = true
	}
}

func wires() {
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
		if red != 0 {
			fmt.Println(cutSecond)
		} else if string(wires[wireLength-1]) == "w" {
			fmt.Println(cutLast)
		} else if blue > 1 {
			fmt.Printf("Cut last blue wire. (The %s wire.)\n", position[lastBlue])
		} else {
			fmt.Println(cutLast)
		}
	case 4:
		recieveSerial()
		if serialOdd && red > 1 {
			fmt.Printf("Cut the last red wire. (The %s wire.)\n", position[lastRed])
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
					fmt.Printf("The value %v in the %s position is incorrect. type again: ", letter, position[i])
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

func removeDupes(list []string) []string {
	var rList []string
	for _, val := range list {
		firstInstance := true
		for _, tVal := range rList {
			if val == tVal {
				firstInstance = false
				break
			}
		}
		if firstInstance {
			rList = append(rList, val)

		}
	}
	return rList
}

func wordGame() {
	words := []string{
		"about", "after", "again", "below", "could",
		"every", "first", "found", "great", "house",
		"large", "learn", "never", "other", "place",
		"plant", "point", "right", "small", "sound",
		"spell", "still", "study", "their", "there",
		"these", "thing", "think", "three", "water",
		"where", "which", "world", "would", "write",
	}
	for i := 0; i < wordLength && len(words) > 1; i++ {
		var letters string
		possibles := []string{}
		fmt.Printf("Please enter the letters in the %s position: ", position[i])
		letters = inputString()
		for j := 0; j < len(words); j++ {
			for k := 0; k < len(letters); k++ {
				if words[j][i] == letters[k] {
					possibles = append(possibles, words[j])
				}
			}
		}
		possibles = removeDupes(possibles)
		if len(possibles) > 1 {
			fmt.Printf("Possible words are: ")
			for i := 0; i < len(possibles); i++ {
				if i == len(possibles)-1 {
					fmt.Printf("%s.", possibles[i])
				} else {
					fmt.Printf("%s, ", possibles[i])
				}
			}
			fmt.Printf("\n")
		}
		words = possibles
	}
	if len(words) == 0 {
		fmt.Printf("There are no words left. Do you want to try again? (yes, no): ")
		replay := yesOrNo()
		if replay == "yes" {
			wordGame()
		}

	} else {
		fmt.Println("Your word is:", words[0])
	}
}

func complexWires() {
	order := []bool{}
	cWires := make(map[string]string)
	cWires = map[string]string{
		"ooo": "c",
		"ool": "d",
		"oso": "c",
		"osl": "b",
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
			fmt.Printf("%s, ", position[index])
		}
	}
	fmt.Printf("positions.\n")
}
