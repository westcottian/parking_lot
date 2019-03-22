package main

import (
	"parking_lot/parkengine"
	"parking_lot/commandsengine"
)

func main() {
	parkengine.NewStorey(4)
	parkengine.NewCar("", "")
	commandsengine.ExecuteFile("samples/file_input.txt")
}
