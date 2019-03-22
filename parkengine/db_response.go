package parkengine

import (
	"fmt"
)

type DbResponse struct {
	storey  Storey
	command string
}

// String returns string output for the logger.
func (d DbResponse) String() string {
	return fmt.Sprintf("Created a parking lot with %d slots", d.storey.maxSlots)
}

// NewDbResponse return an DbResponse Object.
func NewDbResponse(storey Storey, command string) *DbResponse {
	return &DbResponse{storey, command}
}
