package main

import (
	"bufio"
	"os"

	"github.com/BenTaylor0812/ktanbu"
)

func main() {
	ktanbu.Reader = bufio.NewReader(os.Stdin)
	ktanbu.Maze()
	// ktanbu.Position = []string{"first", "second", "third", "fourth", "fifth"}
	// for gameOver := false; !gameOver; {
	// 	fmt.Printf("Enter the game you are playing: ")
	// 	game, _ := ktanbu.Reader.ReadString('\n')
	// 	game = strings.Replace(game, "\r\n", "", -1)
	// 	switch game {
	// 	case "wires", "wire":
	// 		ktanbu.Wires()
	// 	case "button":
	// 		ktanbu.Button()
	// 	case "word", "words":
	// 		ktanbu.Password()
	// 	case "cwires", "cwire", "complexwires", "complex wires":
	// 		ktanbu.ComplexWires()
	// 	case "simon", "simonsays", "simon says":
	// 		ktanbu.Simon()
	// 	case "mazes", "maze":
	// 		ktanbu.MakeNode()
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
