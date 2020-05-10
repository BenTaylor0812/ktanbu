package ktanbu

import "fmt"

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

// Password - Runs the password module
func Password() {
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
		fmt.Printf("Please enter the letters in the %s position: ", Position[i])
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
			Password()
		} else {
			fmt.Println("Okay then.")
		}

	} else {
		fmt.Println("Your word is:", words[0])
	}
}
