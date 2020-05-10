package ktanbu

import (
	"bufio"
)

// Reader - Necessary for reading the inputs for the game.
var Reader *bufio.Reader

// Position - Describes the position with words instead of numbers.
var Position []string

var simonTable [2][3][4]string
var strikes int
var serial string
var serialOdd bool
var serialEven bool
var serialVowel bool
var batteries int
var batteryLock bool
var parallelPort bool
var parallelLock bool
