package models

var (
	// CmdPark command for parking car
	CmdPark = "park"
	// CmdCreateParkingLot command for creating parking lot
	CmdCreateParkingLot = "create_parking_lot"
	// CmdStatus command to finding status of parking lot
	CmdStatus = "status"
	// CmdLeave command for leaving a car
	CmdLeave = "leave"
	// CmdRegistrationNumberByColor command to getting reg number of cars with color
	CmdRegistrationNumberByColor = "registration_numbers_for_cars_with_colour"
	// CmdSlotnoByCarColor command to get slot number of car with color
	CmdSlotnoByCarColor = "slot_numbers_for_cars_with_colour"
	// CmdSlotnoByRegNumber command to get slot number of car with reg number.
	CmdSlotnoByRegNumber = "slot_number_for_registration_number"
)
