package commandsengine

import (
	"fmt"
	assert "github.com/stretchr/testify/assert"
	"testing"
)

func ExampleExecutableFile() {
	ExecuteFile("samples/file_input.txt")
	fmt.Println("")
}

func TestParseCommand(t *testing.T) {
	assert := assert.New(t)

	assert.Equal([]string{"create_parking_lot", "6"}, parseCommand("create_parking_lot   6"))
	assert.Equal([]string{"park", "KA-01-HH-1234", "White"}, parseCommand("park   KA-01-HH-1234	White"))

	assert.True(true)
}
