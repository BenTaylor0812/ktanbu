package ktanbu

import "fmt"

func findButtonColour() string {
	fmt.Printf("Enter the colour of the button: ")
	return inputString()
}

func findButtonText(buttonText *string) string {
	if *buttonText == "" {
		fmt.Printf("Enter the text on the button: ")
		colour := inputString()
		switch colour {
		case "a", "abort":
			*buttonText = "abort"
		case "d", "detonate":
			*buttonText = "detonate"
		case "h", "hold":
			*buttonText = "hold"
		}
	}
	return *buttonText
}

func findLabelIndicator(label string, labelInd *bool) bool {
	fmt.Printf("Is there a lit indicator with the label '%s'?: ", label)
	if yesOrNo() == "yes" {
		*labelInd = true
	} else {
		*labelInd = false
	}
	return *labelInd
}

func holdButton() {
	fmt.Println("Press and hold the button. There should be a band that lights up.")
	fmt.Printf("What colour is it?: ")
	bandColour := inputString()
	switch bandColour {
	case "blue", "b":
		fmt.Println("Release when the timer has a 4 in any position.")
	case "yellow", "y":
		fmt.Println("Release when the timer has a 5 in any position.")
	default:
		fmt.Println("Release when the timer has a 1 in any position.")
	}
}

func releaseButton() {
	fmt.Println("Press and immediately release the button.")
}

// Button - Runs the button modules
func Button() {
	buttonColour := findButtonColour()
	var buttonText string
	var carInd bool
	var frkInd bool
	if buttonColour == "blue" && findButtonText(&buttonText) == "abort" {
		holdButton()
	} else if recieveBatteries() > 1 && findButtonText(&buttonText) == "detonate" {
		releaseButton()
	} else if buttonColour == "white" && findLabelIndicator("CAR", &carInd) {
		holdButton()
	} else if recieveBatteries() > 2 && findLabelIndicator("FRK", &frkInd) {
		releaseButton()
	} else if buttonColour == "yellow" {
		holdButton()
	} else if buttonColour == "red" && findButtonText(&buttonText) == "hold" {
		releaseButton()
	} else {
		holdButton()
	}
}
