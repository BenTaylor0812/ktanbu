package ktanbu

import (
	"fmt"
	"strconv"

	"github.com/BenTaylor0812/paths"
)

func convertToArray(maze string) [rowSize][colSize]string {
	var row int
	var col int
	var mazeMtrx [rowSize][colSize]string
	for _, i := range maze {
		if i == '\n' {
			row++
			col = 0
		} else {
			mazeMtrx[row][col] = string(i)
			col++
		}
	}
	return mazeMtrx
}

func obtainIndex(coord string) int {
	return int(6*(coord[1]-49) + (coord[0] - 97))
}

func makeGraph(mazeNo int) [](*paths.Node) {
	var maze string
	switch mazeNo {
	case 1:
		maze = maze1
	case 2:
		maze = maze2
	case 3:
		maze = maze3
	case 4:
		maze = maze4
	case 5:
		maze = maze5
	case 6:
		maze = maze6
	case 7:
		maze = maze7
	case 8:
		maze = maze8
	case 9:
		maze = maze9
	default:
	}
	mazeMtrx := convertToArray(maze)
	letterMap := map[int]string{
		0: "A",
		1: "B",
		2: "C",
		3: "D",
		4: "E",
		5: "F",
	}
	var graph [](*paths.Node)
	for i, rowArray := range mazeMtrx {
		for j, val := range rowArray {
			var node paths.Node
			node.NodeList = make(map[string]int)
			node.Label = letterMap[j] + strconv.Itoa(i+1)

			switch val {
			case "r":
				node1 := letterMap[j+1] + strconv.Itoa(i+1)
				node2 := letterMap[j] + strconv.Itoa(i+2)
				node.NodeList[node1] = 1
				node.NodeList[node2] = 1
			case "7":
				node1 := letterMap[j] + strconv.Itoa(i+2)
				node2 := letterMap[j-1] + strconv.Itoa(i+1)
				node.NodeList[node1] = 1
				node.NodeList[node2] = 1
			case "J":
				node1 := letterMap[j] + strconv.Itoa(i)
				node2 := letterMap[j-1] + strconv.Itoa(i+1)
				node.NodeList[node1] = 1
				node.NodeList[node2] = 1
			case "L":
				node1 := letterMap[j] + strconv.Itoa(i)
				node2 := letterMap[j+1] + strconv.Itoa(i+1)
				node.NodeList[node1] = 1
				node.NodeList[node2] = 1

			case "-":
				node1 := letterMap[j-1] + strconv.Itoa(i+1)
				node2 := letterMap[j+1] + strconv.Itoa(i+1)
				node.NodeList[node1] = 1
				node.NodeList[node2] = 1
			case "I":
				node1 := letterMap[j] + strconv.Itoa(i)
				node2 := letterMap[j] + strconv.Itoa(i+2)
				node.NodeList[node1] = 1
				node.NodeList[node2] = 1

			case "T":
				node1 := letterMap[j+1] + strconv.Itoa(i+1)
				node2 := letterMap[j] + strconv.Itoa(i+2)
				node3 := letterMap[j-1] + strconv.Itoa(i+1)
				node.NodeList[node1] = 1
				node.NodeList[node2] = 1
				node.NodeList[node3] = 1
			case "K":
				node1 := letterMap[j] + strconv.Itoa(i)
				node2 := letterMap[j+1] + strconv.Itoa(i+1)
				node3 := letterMap[j] + strconv.Itoa(i+2)
				node.NodeList[node1] = 1
				node.NodeList[node2] = 1
				node.NodeList[node3] = 1
			case "_":
				node1 := letterMap[j] + strconv.Itoa(i)
				node2 := letterMap[j+1] + strconv.Itoa(i+1)
				node3 := letterMap[j-1] + strconv.Itoa(i+1)
				node.NodeList[node1] = 1
				node.NodeList[node2] = 1
				node.NodeList[node3] = 1
			case "H":
				node1 := letterMap[j] + strconv.Itoa(i)
				node2 := letterMap[j] + strconv.Itoa(i+2)
				node3 := letterMap[j-1] + strconv.Itoa(i+1)
				node.NodeList[node1] = 1
				node.NodeList[node2] = 1
				node.NodeList[node3] = 1

			case ",":
				node1 := letterMap[j-1] + strconv.Itoa(i+1)
				node.NodeList[node1] = 1
			case ".":
				node1 := letterMap[j+1] + strconv.Itoa(i+1)
				node.NodeList[node1] = 1
			case "!":
				node1 := letterMap[j] + strconv.Itoa(i)
				node.NodeList[node1] = 1
			case "i":
				node1 := letterMap[j] + strconv.Itoa(i+2)
				node.NodeList[node1] = 1
			default:
			}

			graph = append(graph, &node)
		}
	}

	return graph
}

// Maze - function for the maze game
func Maze() {
	var graphNo int
	fmt.Printf("Enter the coords of the two circles: ")
	switch inputString() {
	case "a2 f3", "f3 a2":
		graphNo = 1
	case "b4 e2", "e2 b4":
		graphNo = 2
	case "d4 f4", "f4 d4":
		graphNo = 3
	case "a1 a4", "a4 a1":
		graphNo = 4
	case "d6 e3", "e3 d6":
		graphNo = 5
	case "c5 e1", "e1 c5":
		graphNo = 6
	case "b1 b6", "b6 b1":
		graphNo = 7
	case "d1 c4", "c4 d1":
		graphNo = 8
	case "a5 c2", "c2 a5":
		graphNo = 9
	}

	graph := makeGraph(graphNo)
	fmt.Println(graphNo)
	fmt.Printf("Starting coordinate: ")
	start := inputString()
	fmt.Printf("End coordinate: ")
	end := inputString()
	startIndex := obtainIndex(start)
	endIndex := obtainIndex(end)
	graphArray, _ := paths.ShortestPath(graph[startIndex], graph[endIndex], graph)
	fmt.Println(graphArray)
}
