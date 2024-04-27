package command

import (
	"errors"

	"github.com/codecrafters-io/redis-starter-go/app/resp/data"
)

type Get struct {
	key string
}

var _ Command = &Get{}

func NewGet() Command {
	return &Get{}
}

func (g *Get) Name() string {
	return "GET"
}

func (g *Get) UnmarshalData(request *data.Array) error {
	if key, present := request.StringAt(1); present {
		g.key = key
	} else {
		return errors.New("key not found in command")
	}

	return nil
}

func (g *Get) Key() string {
	return g.key
}

func init() {
	registerCommand(func() Command {
		return NewGet()
	})
}
