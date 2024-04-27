package command

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/resp/data"
)

func UnmarshalBinary(dataBytes []byte) (Command, error) {
	d, err := data.UnmarshalBinary(bytes.NewReader(dataBytes))
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	dataArray, isArray := d.(*data.Array)
	if !isArray {
		return nil, fmt.Errorf("expected command to be Array, got: %T", d)
	}

	if len(dataArray.Elements()) < 1 {
		return nil, fmt.Errorf("expected command array to include at least 1 element, got %d elements", len(dataArray.Elements()))
	}

	commandName, present := dataArray.StringAt(0)
	if !present {
		return nil, errors.New("no command name found")
	}

	command, err := GetCommand(commandName)
	if err != nil {
		return nil, fmt.Errorf("failed to get command: %v", err)
	}

	err = command.UnmarshalData(dataArray)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal command: %w", err)
	}

	return command, nil
}
