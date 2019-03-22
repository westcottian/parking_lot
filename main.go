package main

import (
	"flag"
	"parking_lot/commandsengine"
)

func main() {
	flag.Parse()

	if len(flag.Args()) > 0 {
		commandsengine.ExecuteFile(flag.Args()[0])
		return
	}

	commandsengine.InteractiveSession()
}

