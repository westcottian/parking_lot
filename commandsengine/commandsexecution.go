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
				fmt.Println("first command needs to be creating the storey. Rerun!!!")
				return nil
			}
			maxSlots, err := strToInt(command[1])
			if err != nil {
				fmt.Println(err.Error())
				return err
			}
			// convert this to a new storey addition or update max slot method
			db = parkengine.NewStoreyRunTimeDB(maxSlots)
//			fmt.Println(parkengine.NewDbResponse(*db.Storeys[0], parkengine.CmdCreateParkingLot))
			firstLine = false
			continue
		}
		fmt.Println(processCommand(db, parseCommand(scanner.Text())))	
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
	case parkengine.CmdCreateParkingLot:
		maxSlots, err := strToInt(command[1])
		if err != nil {
			fmt.Println(err.Error())
		}
		return db.AddStorey(maxSlots)
	case parkengine.CmdPark:
		return db.Park(command[1], command[2])
	case parkengine.CmdCreateParkingLot:
	case parkengine.CmdStatus:
		return db.All()
	case parkengine.CmdLeave:
		slotPosition, err := strToInt(command[1])
		if err != nil {
			fmt.Println(err.Error())
		}
		return db.LeaveByPosition(slotPosition)
	case parkengine.CmdRegistrationNumberByColor:
		return db.FindAllByColor(command[1], parkengine.CmdRegistrationNumberByColor)
	case parkengine.CmdSlotnoByCarColor:
		return db.FindAllByColor(command[1], parkengine.CmdSlotnoByCarColor)
	case parkengine.CmdSlotnoByRegNumber:
		return db.FindByRegistrationNumber(command[1])
	default:
	}

	return parkengine.StoreyResponse{}, nil
}

// strToInt conver string to integer
func strToInt(str string) (int, error) {
	nonFractionalPart := strings.Split(str, ".")
	return strconv.Atoi(nonFractionalPart[0])
}

// InteractiveSession take the user through interactive session.
func InteractiveSession() error {
	command := "Start"
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nInput")
	text, _ := reader.ReadString('\n')
	text = strings.TrimRight(text, "\r\n")
	commands := parseCommand(text)
	if text == "" {
		fmt.Println("first command needs to be creating the storey. Rerun!!!")
		return nil
	}
	if commands[0] != parkengine.CmdCreateParkingLot {
		fmt.Println("first command needs to be creating the storey. Rerun!!!")
		return nil
	}
	maxSlots, err := strToInt(commands[1])
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	// convert this to a new storey addition or update max slot method
	db := parkengine.NewStoreyRunTimeDB(maxSlots)
	fmt.Println("\nOutput")
	fmt.Println(parkengine.NewDbResponse(*db.Storeys[0], parkengine.CmdCreateParkingLot))

	for command != "Exit" {
		fmt.Println("\nInput")
		text, _ := reader.ReadString('\n')
		text = strings.TrimRight(text, "\r\n")
		commands := parseCommand(text)
		response, err := processCommand(db, commands)
		if err != nil {
		  	fmt.Println("Error processing output commands.") 
		}
		fmt.Println("\nOutput")
		fmt.Println(response)
		command = commands[0]
	}
	return nil
}
