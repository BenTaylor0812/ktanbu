package ktanbu

import (
	"bufio"
)

// Reader - Necessary for reading the inputs for the game.
var Reader *bufio.Reader

// Position - Describes the position with words instead of numbers.
var Position []string

var serial string
var serialOdd bool
var serialEven bool
var batteries int
var batteryLock bool
var parallelPort bool
var parallelLock bool
