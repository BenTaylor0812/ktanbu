package main

import (
	"bufio"
	"fmt"
	"os"
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

var position []string
var reader *bufio.Reader
var serial string
var serialOdd bool
var batteries int
var parallelPort bool

func main() {
	position = []string{"first", "second", "third", "fourth", "fifth"}
	reader = bufio.NewReader(os.Stdin)
	for gameOver := 0; gameOver != 1; {
		fmt.Printf("Enter the game you are playing: ")
		game, _ := reader.ReadString('\n')
		game = strings.Replace(game, "\r\n", "", -1)
		switch game {
		case "word":
			words()
		case "wires":
			wires()
		case "win":
			fmt.Println("Well done!")
		case "loss":
			fmt.Println("Ah sorry, maybe next time!")
		default:
			fmt.Println("Game over.")
			gameOver = 1
			break
		}
	}
}

func inputString() string {
	output, _ := reader.ReadString('\n')
	return strings.Replace(output, "\r\n", "", -1)
}

func recieveSerial() {
	for serial == "" {
		fmt.Printf("Enter the serial number: ")
		serial = inputString()
		serialOdd = int(serial[len(serial)-1])%2 == 1
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

func words() {
	words := []string{
		"about", "after", "again", "below", "could",
		"every", "first", "found", "great", "house",
		"large", "learn", "never", "other", "place",
		"plant", "point", "right", "small", "sound",
		"spell", "still", "study", "their", "there",
		"these", "thing", "think", "three", "water",
		"where", "which", "world", "would", "write",
	}
	for i := 0; i < wordLength && len(words) != 1; i++ {
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
		if len(possibles) != 1 {
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
	fmt.Println("Your word is:", words[0])
}
