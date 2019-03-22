package models

import "fmt"

type StoreyResponse struct {
	slots   []Slot
	command string
}

func (s StoreyResponse) String() string {
	switch s.command {
	case CmdPark:
		return fmt.Sprintf("Allocated slot number %d", s.slots[0].Position())
	case CmdCreateParkingLot:
	case CmdStatus:
	case CmdLeave:
		return fmt.Sprintf("Slot number %d is free", s.slots[0].Position())
	case CmdRegistrationNumberByColor:
	case CmdSlotnoByCarColor:
	case CmdSlotnoByRegNumber:
	}
	return ""
}
