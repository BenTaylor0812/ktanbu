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

var position []string
var reader *bufio.Reader

func main() {
	position = []string{"first", "second", "third", "fourth", "fifth"}
	reader = bufio.NewReader(os.Stdin)
	// wires()
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

func wires() {
	wires := obtainComb()
	wireLength := len(wires)
	switch wireLength {
	case 3:
	case 4:
	case 5:
	case 6:
	}
}

func obtainComb() string {
	var wires string
	fmt.Printf("Enter the combinations of wires (reading down): ")
	for {
		cVal := 0
		wires, _ = reader.ReadString('\n')
		wires = strings.Replace(wires, "\r\n", "", -1)
		if len(wires) > 6 {
			fmt.Printf("That is too long, type again: ")
			continue
		}
		for i := 0; i < len(wires); i++ {
			letter := string(wires[i])
			switch letter {
			case "k", "w", "b", "r", "y":
				continue
			default:
				fmt.Printf("The value %v in the %s position is incorrect. type again: ", letter, position[i])
				cVal = 1
			}
		}
		if cVal == 0 {
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
		letters, _ = reader.ReadString('\n')
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
