package parkengine

import (
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func TestNewStorey(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	assert.NotEmpty(storey)

	assert.True(true)
}

func TestStorey_FindByRegistrationNumber(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	storey.Park("numberPlate", "color")

	sc := NewSlot(NewCar("numberPlate - x", "color - x"), 0)
	storey.slotList.AddNext(sc)
	// updates the slot position to 3. the current slot order is 1, 3.
	// So next slot should be entered at psoition 2.
	storey.slotList.nextSlot.UpdatePosition(3)
	assert.Equal(3, storey.slotList.nextSlot.Position())
	scy := NewSlot(NewCar("numberPlate - y", "color - y"), 0)
	storey.slotList.AddNext(scy)

	slot, err := storey.FindByRegistrationNumber("numberPlate - y")
	assert.Equal(2, slot.position)
	assert.NoError(err)

	slot, err = storey.FindByRegistrationNumber("numberPlate - z")
	assert.Equal(0, slot.position)
	assert.Error(err)

	slot, err = storey.FindByRegistrationNumber("numberPlate")
	assert.Equal(1, slot.position)
	assert.NoError(err)

	assert.True(true)
}

func TestStorey_FindAllByColor(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	storey.Park("numberPlate", "color")

	sc := NewSlot(NewCar("numberPlate - x", "color"), 0)
	storey.slotList.AddNext(sc)
	// updates the slot position to 3. the current slot order is 1, 3.
	// So next slot should be entered at psoition 2.
	storey.slotList.nextSlot.UpdatePosition(3)
	assert.Equal(3, storey.slotList.nextSlot.Position())
	scy := NewSlot(NewCar("numberPlate - y", "color - y"), 0)
	storey.slotList.AddNext(scy)

	// 2 cars with color present
	slots, err := storey.FindAllByColor("color")
	assert.Equal(2, len(slots))
	assert.Equal(3, slots[0].Position())
	assert.Equal(1, slots[1].Position())
	assert.NoError(err)

	// no cars with color present
	slots, err = storey.FindAllByColor("color - z")
	assert.Equal(0, len(slots))
	assert.Error(err)

	// 1 car with color present
	slots, err = storey.FindAllByColor("color - y")
	assert.Equal(1, len(slots))
	assert.Equal(2, slots[0].Position())
	assert.NoError(err)

	assert.True(true)
}

func TestStorey_OccupancyCount(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	assert.Equal(0, storey.OccupancyCount())

	storey.Park("numberPlate", "color")
	assert.Equal(1, storey.OccupancyCount())

	assert.True(true)
}

func TestStorey_Park(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	assert.Equal(0, storey.OccupancyCount())

	storey.Park("numberPlate", "color")
	assert.Equal(1, storey.OccupancyCount())

	storey.slotList.UpdatePosition(2)
	storey.Park("numberPlate - x", "color - x")

	assert.Equal("numberPlate - x", storey.slotList.car.numberPlate)
	assert.Equal(1, storey.slotList.Position())

	assert.True(true)
}

func TestSlot_AddNext(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	assert.Equal(0, storey.OccupancyCount())

	storey.Park("numberPlate", "color")
	assert.Equal(1, storey.OccupancyCount())

	sc := NewSlot(NewCar("numberPlate - x", "color - x"), 0)
	storey.slotList.AddNext(sc)

	assert.Equal(1, storey.slotList.Position())
	assert.Equal(2, storey.OccupancyCount())
	assert.Equal(2, storey.slotList.nextSlot.Position())

	// updates the slot position to 3. the current slot order is 1, 3.
	// So next slot should be entered at psoition 2.
	storey.slotList.nextSlot.UpdatePosition(3)
	assert.Equal(3, storey.slotList.nextSlot.Position())

	scy := NewSlot(NewCar("numberPlate - y", "color - y"), 0)
	storey.slotList.AddNext(scy)
	assert.Equal(2, storey.slotList.nextSlot.Position())
	assert.Equal(3, storey.slotList.nextSlot.nextSlot.Position())

	assert.Equal("numberPlate - y", storey.slotList.nextSlot.car.numberPlate)
	assert.Equal("numberPlate - x", storey.slotList.nextSlot.nextSlot.car.numberPlate)

	assert.True(true)
}

func TestStorey_Leave(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	storey.Park("numberPlate", "color")

	sc := NewSlot(NewCar("numberPlate - x", "color - x"), 0)
	storey.slotList.AddNext(sc)
	// updates the slot position to 3. the current slot order is 1, 3.
	// So next slot should be entered at psoition 2.
	storey.slotList.nextSlot.UpdatePosition(3)
	assert.Equal(3, storey.slotList.nextSlot.Position())
	scy := NewSlot(NewCar("numberPlate - y", "color - y"), 0)
	storey.slotList.AddNext(scy)

	storey.Leave("numberPlate")

	assert.Equal("numberPlate - y", storey.slotList.car.numberPlate)
	assert.Equal("numberPlate - x", storey.slotList.nextSlot.car.numberPlate)
	assert.Equal(2, storey.slotList.Position())
	assert.Equal(3, storey.slotList.nextSlot.Position())

	assert.True(true)
}

func TestSlot_Leave(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	storey.Park("numberPlate", "color")

	sc := NewSlot(NewCar("numberPlate - x", "color - x"), 0)
	storey.slotList.AddNext(sc)
	// updates the slot position to 3. the current slot order is 1, 3.
	// So next slot should be entered at psoition 2.
	storey.slotList.nextSlot.UpdatePosition(3)
	assert.Equal(3, storey.slotList.nextSlot.Position())

	scy := NewSlot(NewCar("numberPlate - y", "color - y"), 0)
	storey.slotList.AddNext(scy)

	storey.slotList.nextSlot.Leave()

	assert.Equal("numberPlate", storey.slotList.car.numberPlate)
	assert.Equal("numberPlate - x", storey.slotList.nextSlot.car.numberPlate)
	assert.Equal(1, storey.slotList.Position())
	assert.Equal(3, storey.slotList.nextSlot.Position())

	assert.True(true)
}

func TestSlot_AddNext_PrevSlot(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	storey.Park("numberPlate", "color")

	sc := NewSlot(NewCar("numberPlate - x", "color - x"), 0)
	storey.slotList.AddNext(sc)
	// updates the slot position to 3. the current slot order is 1, 3.
	// So next slot should be entered at psoition 2.
	storey.slotList.nextSlot.UpdatePosition(3)
	assert.Equal(3, storey.slotList.nextSlot.Position())

	scy := NewSlot(NewCar("numberPlate - y", "color - y"), 0)
	storey.slotList.AddNext(scy)

	latestSlot := storey.slotList.nextSlot.nextSlot

	assert.Equal("numberPlate - y", latestSlot.prevSlot.car.numberPlate)
	assert.Equal("numberPlate - x", latestSlot.car.numberPlate)
	assert.Equal("numberPlate", latestSlot.prevSlot.prevSlot.car.numberPlate)
	assert.Equal(1, latestSlot.prevSlot.prevSlot.Position())
	assert.Equal(2, latestSlot.prevSlot.Position())
	assert.Equal(3, latestSlot.Position())

	assert.True(true)
}

func TestSlot_FindCar(t *testing.T) {
	assert := assert.New(t)

	storey := NewStorey(4)
	storey.Park("numberPlate", "color")

	s2, err := storey.Park("numberPlate - x", "color - x")
	assert.Equal(2, s2.Position())

	s3, err := storey.Park("numberPlate - y", "color - y")
	assert.NoError(err)
	assert.Equal(3, s3.Position())
	assert.Equal(s3.position, s2.nextSlot.Position())

	assert.Equal(2, storey.slotList.nextSlot.Position())
	assert.Equal(3, storey.slotList.nextSlot.nextSlot.Position())

	sct, err := storey.slotList.FindCar("numberPlate - y")
	assert.NoError(err)
	assert.Equal(3, sct.Position())

	sct, err = storey.slotList.FindCar("numberPlate - x")
	assert.NoError(err)
	assert.Equal(2, sct.Position())

	sct, err = storey.slotList.FindCar("numberPlate - z")
	assert.Error(err)
	assert.True(true)
}

func TestNewCar(t *testing.T) {
	assert := assert.New(t)

	car := NewCar("KL-00-0000", "red")
	assert.NotEmpty(car)

	assert.True(true)
}
