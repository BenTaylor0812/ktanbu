package ktanbu

import (
	"fmt"
	"strconv"
	"strings"
)

func yesOrNo() string {
	var response string
	for {
		response = inputString()
		switch response {
		case "yes", "y":
			return "yes"
		case "no", "n":
			return "no"
		default:
			fmt.Printf("Please enter yes or no: ")
			continue
		}
	}
}

func inputString() string {
	output, _ := Reader.ReadString('\n')
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
	}
	serialOdd = int(serial[len(serial)-1])%2 == 1
	serialEven = !serialOdd
	for _, i := range serial {
		switch string(i) {
		case "a", "e", "i", "o", "u":
			serialVowel = true
		}
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

func recieveBatteries() int {
	for !batteryLock {
		fmt.Printf("How many batteries are there? ")
		batteries = inputInteger()
		batteryLock = true
	}
	return batteries
}
