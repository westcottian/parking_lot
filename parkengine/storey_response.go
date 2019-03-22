package parkengine

import (
	"fmt"
	"strconv"
	"strings"
)

type StoreyResponse struct {
	slots   []Slot
	storey  *Storey
	command string
}

func (s StoreyResponse) String() string {
	switch s.command {
	case CmdCreateParkingLot:
		return fmt.Sprintf("Created a parking lot with %d slots", s.storey.maxSlots)
	case CmdPark:
		return fmt.Sprintf("Allocated slot number: %d", s.slots[0].Position())
	case CmdCreateParkingLot:
	case CmdStatus:
		content := fmt.Sprintf("Slot No.\tRegistration No\tColor")
		for _, slot := range s.slots {
			content += fmt.Sprintf("\n%d %s %s", slot.Position(), slot.RegistrationNumber(), slot.Color())
		}
		return content
	case CmdLeave:
		return fmt.Sprintf("Slot number %d is free", s.slots[0].Position())
	case CmdRegistrationNumberByColor:
		regNumbers := []string{}
		for _, s := range s.slots {
			regNumbers = append(regNumbers, s.car.numberPlate)
		}

		return strings.Join(regNumbers, ", ")
	case CmdSlotnoByCarColor:
		positions := []string{}
		for _, s := range s.slots {
			positions = append(positions, strconv.Itoa(s.Position()))
		}

		return strings.Join(positions, ", ")
	case CmdSlotnoByRegNumber:
		return strconv.Itoa(s.slots[0].Position())
	}
	return ""
}
