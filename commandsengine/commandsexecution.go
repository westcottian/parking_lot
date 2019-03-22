package commandsengine

import (
	"bufio"
	"log"
	"os"
	"strings"
	"parking_lot/parkengine"
	"fmt"
	"strconv"
)

var (
	// CommandSeparator is the separator used by default in the input
	CommandSeparator = " "
	// Tab is the tab character.
	Tab = "\t"
)

// ExecuteFile taskes in a file path and execute the commands in the file.
func ExecuteFile(filepath string) error {
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	db := parkengine.NewStoreyRunTimeDB(0)

	firstLine := true
	for scanner.Scan() {
		if firstLine {
			text := scanner.Text()
			command := parseCommand(text)
			if command[0] != parkengine.CmdCreateParkingLot {
				panic("first command needs to be creating the storey")
			}
			maxSlots, err := strToInt(command[1])
			if err != nil {
				panic(err.Error())
			}
			// convert this to a new storey addition or update max slot method
			db = parkengine.NewStoreyRunTimeDB(maxSlots)
			firstLine = false
			continue
		}
		fmt.Println(db)	
	}	


	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}

// parseCommand takes in a command string and converts it to a string array
// by removing the tab character in the command.
func parseCommand(command string) []string {
	parsedCommand := []string{}

	// remove the tabs in between the command.
	command = strings.Replace(command, Tab, CommandSeparator, -1)

	// remove the empty string
	for _, s := range strings.Split(command, CommandSeparator) {
		if s != "" {
			parsedCommand = append(parsedCommand, s)
		}
	}

	return parsedCommand
}

// processCommand process each command
func processCommand(db parkengine.DataStore, command []string) (parkengine.StoreyResponse, error) {
	switch command[0] {
	case parkengine.CmdPark:
		return db.Park(command[1], command[2])
	case parkengine.CmdCreateParkingLot:
	case parkengine.CmdStatus:
	case parkengine.CmdLeave:
		slotPosition, err := strToInt(command[1])
		if err != nil {
			panic(err.Error())
		}
		return db.LeaveByPosition(slotPosition)
	case parkengine.CmdRegistrationNumberByColor:
	case parkengine.CmdSlotnoByCarColor:
	case parkengine.CmdSlotnoByRegNumber:
	default:
	}

	return parkengine.StoreyResponse{}, nil
}

// strToInt conver string to integer
func strToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}
