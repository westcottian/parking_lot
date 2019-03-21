package parkengine

import (
	"errors"
)

var (
	ErrParkingFull= errors.New("Parking Full")
	ErrParkingAvailable = errors.New("Parking Available")
	ErrCarNotParked= errors.New("Car not found")
)

//Slot struct
type Slot struct {
	prevSlot *Slot
	car      *Car
	position int
	nextSlot *Slot
}

type Storey struct {
	maxSlots int
	slotList *Slot   //Form LinkList
}

// Car - define the car preoperties
type Car struct {
	numberPlate string
	color       string
}

// CarPark - check if the Slot is available, if available Create Slot in the vacancy and associate with adjacent slots
// return Slot
func (s *Storey) Park(numberPlate, color string) (*Slot, error) {
	slot := &Slot{}
	if s.OccupancyCount() >= s.maxSlots {
		return slot, ErrParkingFull
	}

	car := NewCar(numberPlate, color)

	if s.OccupancyCount() == 0 {
		slot := NewSlot(car, 1)
		s.slotList = slot
		return slot, nil
	}

	if s.slotList.Position() > 1 {
		currSlot := s.slotList
		s.slotList = NewSlot(car, 1)
		s.slotList.AddNext(currSlot)
	}

	slot = NewSlot(car, 0)
	s.slotList.AddNext(slot)

	return slot, nil
}

// Leave - check if the Slot is available
// if available Create Slot in the vacancy and associate with adjacent slots
// return Slot
func (s *Storey) Leave(numberPlate string) (*Slot, error) {
	if s.slotList == nil {
		return &Slot{}, ErrParkingAvailable
	}

	return s.slotList.FindCar(numberPlate)
}

// OccupancyCount returns the number of slots occupied in this storey.
func (s *Storey) OccupancyCount() int {
	if s.slotList == nil {
		return 0
	}

	return s.slotList.CountSelf()
}

// NewStorey returns a Storey object
func NewStorey(maxSlots int) *Storey {
	return &Storey{
		maxSlots: maxSlots,
	}
}

// Leave - leave the Car, and connect the prev slot with next
func (s *Slot) Leave() error {
	return nil
}

// FindCar - finds if the slot has the car or else check in the next slot
func (s *Slot) FindCar(numberPlate string) (*Slot, error) {
	if s.car.numberPlate == numberPlate {
		return s, nil
	}

	if s.nextSlot == nil {
		return &Slot{}, ErrCarNotParked
	}

	return s.nextSlot.FindCar(numberPlate)
}

// AddNext - add a new Slot after the current and associate the current next to the new.
func (s *Slot) AddNext(sc *Slot) error {
	if s.nextSlot == nil {
		s.nextSlot = sc.UpdatePosition(s.position + 1)
		return nil
	}

	if s.nextSlot.position > (s.position + 1) {
		currentNext := s.nextSlot
		s.nextSlot = sc.UpdatePosition(s.position + 1)
		sc.nextSlot = currentNext
		return nil
	}

	s.nextSlot.AddNext(sc)

	return nil
}

// CountSelf counts 1 for self and relayes the count Self to next Slot.
func (s Slot) CountSelf() int {
	if s.nextSlot == nil {
		return 1
	}

	return 1 + s.nextSlot.CountSelf()
}

// UpdatePosition updates the position ofthe slot to the specified position value.
func (s *Slot) UpdatePosition(position int) *Slot {
	s.position = position
	return s
}

// Position return the position of the Slot
func (s Slot) Position() int {
	return s.position
}

// NewSlot returns a slot object
func NewSlot(car *Car, position int) *Slot {
	return &Slot{car: car, position: position}
}

// NewCar returns a new car object
func NewCar(numberPlate, color string) *Car {
	return &Car{
		numberPlate: numberPlate,
		color:       color,
	}
}
