package command

import (
	"fmt"

	"github.com/codecrafters-io/redis-starter-go/app/resp/data"
)

type Echo struct {
	data string
}

var _ Command = &Echo{}

func NewEchoCommand() Command {
	return &Echo{}
}

func (c *Echo) Name() string {
	return "ECHO"
}

func (c *Echo) UnmarshalData(request *data.Array) error {
	// Data to be echoed is a bulk string in the second position
	argument := request.Elements()[1]

	if dataToBeEchoed, isBulkString := argument.(*data.BulkString); !isBulkString {
		return fmt.Errorf("expected argument to be bulk string, got %T", argument)
	} else {
		c.data = dataToBeEchoed.Data()
	}

	return nil
}

func (c *Echo) Data() string {
	return c.data
}

func init() {
	registerCommand(func() Command {
		return NewEchoCommand()
	})
}
