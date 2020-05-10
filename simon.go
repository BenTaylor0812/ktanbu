package ktanbu

import "fmt"

func simonBody(index int) {
	simonTable = [2][3][4]string{
		{
			{"b", "r", "y", "g"},
			{"y", "g", "b", "r"},
			{"g", "r", "y", "b"},
		},
		{
			{"b", "y", "g", "r"},
			{"r", "b", "y", "g"},
			{"y", "g", "b", "r"},
		},
	}

	combo := map[string]string{
		"r": simonTable[index][strikes][0],
		"b": simonTable[index][strikes][1],
		"g": simonTable[index][strikes][2],
		"y": simonTable[index][strikes][3],
	}
	fmt.Println(combo)

	var record []string
	var tRecord []string
	for {
		fmt.Printf("Enter the order of flashing lights: ")
		printSlice(record)
		var add string
		var done bool
		add = inputString()
		switch add {
		case "r", "b", "g", "y":
			record = append(record, add)
			tRecord = append(tRecord, combo[add])
			fmt.Printf("The order you have to press is: ")
			printSlice(tRecord)
			inputString()
			fmt.Printf("\n")
		case "d":
			done = true
			break
		case "e":
			strikes++
			if strikes < 3 {
				fmt.Println("An error occured, please restart.")
				simonBody(index)
			} else {
				fmt.Println("Eek you died!")
			}
			done = true
			break
		}
		if done {
			break
		}
	}
}

func printSlice(record []string) {
	for _, i := range record {
		fmt.Printf(i)
	}
}

// Simon - Runs the simon says module
func Simon() {
	fmt.Printf("Enter how many strikes you have: ")
	strikes = inputInteger()
	recieveSerial()
	if serialVowel {
		simonBody(vowelPresent)
	} else {
		simonBody(noVowelPresent)
	}
}
