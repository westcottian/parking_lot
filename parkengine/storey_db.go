package parkengine


// storeyDB holds the data in memory while run time.
// in requirements document multi storey is mentioned.
// but no further actions are requested in the same.
// so the Storey is defined as an array.
type storeyDB struct {
	Storeys []*Storey
}

// NewStoreyRunTimeDB returns an instance of the storey db.
func NewStoreyRunTimeDB(maxSlots int) *storeyDB {
	storey := NewStorey(maxSlots)
	return &storeyDB{
		[]*Storey{
			storey,
		},
	}
}

// Park a car
func (s *storeyDB) Park(numberPlate, color string) (StoreyResponse, error) {
	// until we start supporting more than one storey
	slot, err := s.Storeys[0].Park(numberPlate, color)
	sResponse := StoreyResponse{
		slots: []Slot{
			*slot,
		},
		command: CmdPark,
	}

	return sResponse, err
}

// LeaveByPosition leave a car froma slot by the position
func (s *storeyDB) LeaveByPosition(position int) (StoreyResponse, error) {
	slot, err := s.Storeys[0].LeaveByPosition(position)
	sResponse := StoreyResponse{
		slots: []Slot{
			*slot,
		},
		command: CmdLeave,
	}

	return sResponse, err
}

// FindByRegistrationNumber find slot having car with registration number.
func (s *storeyDB) FindByRegistrationNumber(numberPlate string) (StoreyResponse, error) {
	return StoreyResponse{}, nil
}

// FindAllByColor find the slots having cars with the color.
func (s *storeyDB) FindAllByColor(color string) (StoreyResponse, error) {
	return StoreyResponse{}, nil
}
